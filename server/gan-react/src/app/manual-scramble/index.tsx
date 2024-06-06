import React, { useState, useEffect, useCallback } from "react";
import { WebSocketServer } from "ws";
import { useDispatch, useSelector } from "react-redux";
import { Button, TextField, Tooltip } from "@material-ui/core";
import { FormattedMessage } from "react-intl";

import { fiveSideSolver } from "core/cube/solvers/five-side-solver";
import { validateAlgorithm, invertAlgorithm } from "core/cube/libs/algorithms";
import { ButtonRow, ContentContainer } from "core/components/presentation";

import { getRobotServer } from "app/robot/store/selectors";
import { CubePreview } from "app/cube-preview";
import { scrambleSubmitted } from "app/robot/store/actions";

interface ScrambleState {
  previewCode: string;
  hasError: boolean;
}

export const ManualScramble = (): JSX.Element => {
  const [scramble, setScramble] = useState<ScrambleState>({
    //previewCode: "F2 B2 L2 B2 R2 D2 F2 B2 D2 F2 R2 B2 D2 B2 L2 D2 L2",
    previewCode: "F2 B2",
    hasError: false,
  });
  const robotServer = useSelector(getRobotServer);
  const dispatch = useDispatch();

  // WebSocket state
  const [ws, setWs] = useState<WebSocket | null>(null);

  // WebSocket connection setup
  useEffect(() => {
    const socket = new WebSocket("ws://localhost:3000");
    setWs(socket);

    socket.onopen = () => {
      console.log("Connected to WebSocket server");
    };

    socket.onmessage = (event) => {
      const data = JSON.parse(event.data);
      console.log(data);
      handleBackendMessage(data);
    };

    socket.onclose = () => {
      console.log("Disconnected from WebSocket server");
    };

    return () => {
      socket.close();
    };
  }, []);

  // Handle messages from the WebSocket server
  const handleBackendMessage = (data: any) => {
    if (data.payload === "connectToRobot") {
      console.log("Connected to the robot");
      //connectToRobot();
    } else if (data.payload === "sendInstruction") {
      console.log("Instruction received");
      setScramble({
        previewCode: "",
        hasError: false,
      });
      console.log(scramble.previewCode);
      handleSendClick();
      //processScrambleInput(data.instruction);
    }
  };

  const processScrambleInput = useCallback((scramble: string) => {
    if (scramble.length > 0) {
      if (validateAlgorithm(scramble)) {
        const fiveSideSolve = fiveSideSolver(scramble);
        if (fiveSideSolve) {
          const fiveSideScramble = invertAlgorithm(fiveSideSolve);
          setScramble({
            previewCode: fiveSideScramble,
            hasError: false,
          });
        }
      } else {
        setScramble({
          previewCode: "",
          hasError: true,
        });
      }
    } else {
      setScramble({
        previewCode: "",
        hasError: false,
      });
    }
  }, []);

  const handleManualScrambleChange = useCallback(
    (event: React.FocusEvent<HTMLInputElement>) =>
      processScrambleInput(event.currentTarget.value),
    []
  );

  const handleManualScrambleKeyDown = useCallback((e: React.KeyboardEvent) => {
    const target = e.target as HTMLInputElement;
    if (e.key === "’") {
      target.value += "'";
      e.preventDefault();
      e.stopPropagation();
    }
  }, []);

  const handleSendClick = useCallback(
    () => dispatch(scrambleSubmitted(scramble.previewCode)),
    [scramble]
  );

  const tooltipText = !Boolean(robotServer)
    ? "Robot not connected"
    : !scramble
    ? "Scramble required"
    : "";

  return (
    <ContentContainer>
      <form noValidate autoComplete="off" className="container">
        <TextField
          id="manual-scramble"
          label="Manual Scramble"
          multiline
          rowsMax={4}
          fullWidth
          InputLabelProps={{
            shrink: true,
          }}
          error={scramble.hasError}
          helperText={scramble.hasError && "Invalid Scramble Code"}
          onChange={handleManualScrambleChange}
          onKeyDown={handleManualScrambleKeyDown}
        />
      </form>
      <CubePreview scrambleCode={scramble.previewCode} />
      <ButtonRow>
        <Tooltip arrow title={tooltipText}>
          <span>
            <Button
              variant="contained"
              disabled={!Boolean(robotServer) || !Boolean(scramble)}
              onClick={handleSendClick}
              className="flex flex-grow"
            >
              <FormattedMessage id="scramble.actions.send" />
            </Button>
          </span>
        </Tooltip>
      </ButtonRow>
    </ContentContainer>
  );
};
