#pragma once
#include "Interfaces/IRuntimeFactory.h"
#include "Interfaces/IConfigRuntimeFactory.h"
#include "Interfaces/ITcpConnectionData.h"

namespace JavonetNS::Cpp::Sdk {
	/// \brief The Javonet class is a singleton class that serves as the entry point for interacting with Javonet.
	/// It provides methods to activate and initialize the Javonet SDK.
	/// It supports both in-memory and TCP connections.
	/// Refer to this <a href="https://www.javonet.com/guides/v2/cpp/foundations/javonet-static-class">article on Javonet Guides</a>
	class Javonet {
	private:
		Javonet();
		static int Activate();
	public:
		~Javonet() {};
		static class _init {
		public:
			_init();
		} _initializer;

		/// \brief Activates Javonet with the provided license key.
		/// Refer to this <a href="https://www.javonet.com/guides/v2/cpp/getting-started/activating-javonet">article on Javonet Guides</a>
		/// \param licenseKey The license key to activate Javonet.
		/// \return The activation status code.
		dllExport static int Activate(std::string licenseKey);

		/// \brief Activates Javonet with the provided license key and proxy host.
		/// Refer to this <a href="https://www.javonet.com/guides/v2/cpp/getting-started/activating-javonet">article on Javonet Guides</a>
		/// \param licenseKey The license key to activate Javonet.
		/// \param proxyHost The host for the proxy server.
		/// \param proxyUserName The username for the proxy server (optional).
		/// \param proxyUserPassword The password for the proxy server (optional).
		/// \return The activation status code.
		dllExport static int Activate(std::string licenseKey, std::string proxyHost, std::string proxyUserName, std::string proxyPassword);

		/// \brief Initializes Javonet using an in-memory channel on the same machine.
		/// Refer to this <a href="https://www.javonet.com/guides/v2/cpp/foundations/in-memory-channel">article on Javonet Guides</a>
		/// \return A RuntimeFactory instance configured for an in-memory connection.
		dllExport static std::unique_ptr<JavonetNS::Cpp::Sdk::Interfaces::IRuntimeFactory> InMemory();

		/// \brief Initializes Javonet with a TCP connection to a remote machine.
		/// Refer to this <a href="https://www.javonet.com/guides/v2/cpp/foundations/tcp-channel">article on Javonet Guides</a>
		/// \param tcpConnectionData The address of the remote machine.
		/// \return A RuntimeFactory instance configured for a TCP connection.
		dllExport static std::unique_ptr<JavonetNS::Cpp::Sdk::Interfaces::IRuntimeFactory> Tcp(std::shared_ptr<JavonetNS::Cpp::Utils::Interfaces::ITcpConnectionData> tcpConnectionData);

		/// \brief Initializes Javonet with a configuration file.
		/// Currently supported: Configuration file in JSON format.
		/// Refer to this <a href="https://www.javonet.com/guides/v2/cpp/foundations/configure-channel">article on Javonet Guides</a>
		/// \param path The path to the configuration file.
		/// \return A ConfigureRuntimeFactory instance with configuration data.
		dllExport static std::unique_ptr<JavonetNS::Cpp::Sdk::Interfaces::IConfigRuntimeFactory> WithConfig(std::string path);

	};
}
