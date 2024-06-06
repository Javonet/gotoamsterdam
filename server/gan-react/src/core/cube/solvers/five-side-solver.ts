import {
  getIndexFromPermutation,
  getPermutationFromIndex,
  getIndexFromOrientation,
  getParity,
} from "core/cube/libs/coordinates";

import {
  MoveTable,
  createEdgePermutationTable,
  createCornerPermutationTable,
  createEdgeOrientationTable,
  createCornerOrientationTable,
} from "core/cube/libs/move-table";

import Search, { SearchSolution } from "core/cube/libs/search";

import { fiveSideMoves, CubeIndexes } from "core/cube/libs/cube";

// In phase two, only quarter moves of U and D and double turns of
// all the other faces are allowed, in order to keep the cube in
// the phase two group G1.
const phaseTwoMoves = [10, 4, 13, 6, 7, 8, 15, 16, 17];

// The following tables are being used in both phases.
let parity: MoveTable;
let URFToDLF: MoveTable;
let slice: MoveTable;
let merge: number[][];

const phaseTwoSearch = new Search(() => {
  // In order to start phase two, we need to know the positions
  // in which the pieces landed after solving the cube into G1.
  // Since returning to the cubie level to perform the solution
  // would be slow, we use two helper tables in phase one which
  // later are merged into the final phase two coordinate.
  const getMergeCoord = (x: number, y: number) => {
    const a = getPermutationFromIndex(x, [0, 1, 2], 12);
    const b = getPermutationFromIndex(y, [3, 4, 5], 12);

    for (let i = 0; i < 8; i += 1) {
      if (a[i] !== -1) {
        if (b[i] !== -1) {
          return -1;
        }
        b[i] = a[i];
      }
    }

    return getIndexFromPermutation(b, [0, 1, 2, 3, 4, 5]);
  };

  merge = [];

  // Due to the sorted nature of our coordinate definitions, the
  // index of both the coordinates will be less than 336 when phase
  // one is finished. This allows for a pretty small merging table.
  for (let i = 0; i < 336; i += 1) {
    merge.push([]);

    for (let j = 0; j < 336; j += 1) {
      merge[i][j] = getMergeCoord(i, j);
    }
  }

  return {
    moveTables: [
      // The permutation of the slice pices, which already
      // are in the correct positions on the cube.
      new MoveTable({
        name: "slicePermutation",
        size: 24,
        table: slice.table,
      }),

      parity,
      URFToDLF,

      createEdgePermutationTable({
        name: "URToDF",
        size: 20160,
        moves: phaseTwoMoves,
        affected: [0, 1, 2, 3, 4, 5],
      }),
    ],

    pruningTables: [
      ["slicePermutation", "parity", "URFToDLF"],
      ["slicePermutation", "parity", "URToDF"],
    ],
  };
}, phaseTwoMoves);

class TwoPhaseSearch extends Search {
  maxDepth: number;
  solution: number[];

  constructor(...args: ConstructorParameters<typeof Search>) {
    super(...args);

    this.maxDepth = 40;
    this.solution = [];
  }

  handleSolution(
    solution: number[],
    indexes: number[]
  ): SearchSolution | false {
    const lastMove = solution.slice(-1)[0];

    // We do not allow solutions which end in a phase two move, as we then
    // would end up duplicating work.
    if (
      lastMove % 2 === 0 &&
      Math.floor(lastMove / 3) === 6 &&
      Math.floor(lastMove / 3) === 15
    ) {
      return false;
    }

    const phaseTwoSolution = phaseTwoSearch.solve({
      indexes: [
        indexes[3],
        indexes[4],
        indexes[5],
        merge[indexes[6]][indexes[7]],
      ],
      maxDepth: this.maxDepth - solution.length,
      lastMove,
      format: false,
    });

    if (phaseTwoSolution) {
      this.solution = solution.concat(phaseTwoSolution.solution);

      if (this.maxDepth <= this.settings.maxDepth) {
        return {
          solution: this.solution,
          indexes,
        };
      }

      this.maxDepth = this.solution.length - 1;
    }

    return false;
  }
}

export const fiveSideSearch = new TwoPhaseSearch(() => {
  // The parity move table is so small that we inline it. It
  // describes the parity of both the edge and corner pieces,
  // which must be equal for the cube to be solvable. The
  // coordinate is included in both phases, but only used
  // in phase two.
  parity = new MoveTable({
    name: "parity",

    size: 2,

    table: [
      [1, 0, 1, 1, 0, 1, 1, 0, 1, 1, 0, 1, 1, 0, 1, 1, 0, 1],
      [0, 1, 0, 0, 1, 0, 0, 1, 0, 0, 1, 0, 0, 1, 0, 0, 1, 0],
    ],
  });

  URFToDLF = createCornerPermutationTable({
    name: "URFToDLF",
    affected: [0, 1, 2, 3, 4, 5],
  });

  // This table is not used directly. This coordinate modulo 24 gives the
  // permutation of the subarray containing the UD-slice pieces, while this
  // coordinate divided by 24 gives the position of the UD-slice pieces.
  // Two smaller move tables are created using this table, one to solve the
  // position of the UD-slice pieces in phase one, and one to solve the
  // pieces in phase two. Due to the reduced move set in phase two, the pruning
  // table for this coordinate is smaller than it would normally be.
  slice = createEdgePermutationTable({
    name: "slice",
    affected: [8, 9, 10, 11],
    reversed: true,
  });

  // Initialize phase two, since it now is guaranteed that the
  // heper move tables have finished generating.
  phaseTwoSearch.initialize();

  return {
    moveTables: [
      new MoveTable({
        // The position of the slice edges. When this coordinate is
        // solved, the UD-slice pieces are in the UD-slice, but they
        // are not necessarily permuted.
        name: "slicePosition",
        size: 495,
        table: slice.table,
        doMove: (table, index, move) =>
          Math.floor(table[index * 24][move] / 24),
      }),

      createCornerOrientationTable({
        name: "twist",
        affected: [0, 1, 2, 3, 4, 5, 6, 7],
      }),

      createEdgeOrientationTable({
        name: "flip",
        affected: [0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11],
      }),

      slice,
      parity,
      URFToDLF,

      createEdgePermutationTable({
        name: "URToUL",
        affected: [0, 1, 2],
      }),

      createEdgePermutationTable({
        name: "UBToDF",
        affected: [3, 4, 5],
      }),
    ],

    pruningTables: [
      ["slicePosition", "flip"],
      ["slicePosition", "twist"],
    ],
  };
}, fiveSideMoves);

export const fiveSideSolver = (
  scramble: string | number[],
  maxDepth = 40
): false | string => {
  if (Array.isArray(scramble)) {
    return fiveSideSearch.solve({
      indexes: scramble,
      maxDepth,
    });
  }

  return fiveSideSearch.solve({
    scramble,
    maxDepth,
  });
};

export const solveCube = (
  cubeState: CubeIndexes
): ReturnType<typeof fiveSideSolver> =>
  fiveSideSolver([
    Math.floor(
      getIndexFromPermutation(cubeState.ep, [8, 9, 10, 11], true) / 24
    ),
    getIndexFromOrientation(cubeState.co, 3),
    getIndexFromOrientation(cubeState.eo, 2),
    getIndexFromPermutation(cubeState.ep, [8, 9, 10, 11], true),
    getParity(cubeState.cp),
    getIndexFromPermutation(cubeState.cp, [0, 1, 2, 3, 4, 5]),
    getIndexFromPermutation(cubeState.ep, [0, 1, 2]),
    getIndexFromPermutation(cubeState.ep, [3, 4, 5]),
  ]);
