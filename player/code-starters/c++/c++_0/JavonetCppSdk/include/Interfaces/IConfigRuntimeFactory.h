#pragma once
#include "Interfaces/IRuntimeContext.h"

namespace JavonetNS::Cpp::Sdk::Interfaces {
	/// \brief The IConfigRuntimeFactory interface provides methods for creating runtime contexts.
	/// Each method corresponds to a specific runtime (CLR, JVM, .NET Core, Perl, Ruby, Node.js, Python) and returns a shared_ptr to an IRuntimeContext instance for that runtime.
	class IConfigRuntimeFactory {
	public:
		virtual ~IConfigRuntimeFactory() = default;

		/// \brief Creates IRuntimeContext instance to interact with CLR runtime.
		/// \param configName The name of the configuration to use.
		/// \return A shared_ptr to an IRuntimeContext instance for the CLR runtime.
		virtual std::shared_ptr<IRuntimeContext> Clr(std::string configName = "default") = 0;

		/// \brief Creates IRuntimeContext instance to interact with JVM runtime.
		/// \param configName The name of the configuration to use.
		/// \return A shared_ptr to an IRuntimeContext instance for the JVM runtime.
		virtual std::shared_ptr<IRuntimeContext> Jvm(std::string configName = "default") = 0;

		/// \brief Creates IRuntimeContext instance to interact with .NET runtime.
		/// \param configName The name of the configuration to use.
		/// \return A shared_ptr to an IRuntimeContext instance for the .NET runtime.
		virtual std::shared_ptr<IRuntimeContext> Netcore(std::string configName = "default") = 0;

		/// \brief Creates IRuntimeContext instance to interact with Perl runtime.
		/// \param configName The name of the configuration to use.
		/// \return A shared_ptr to an IRuntimeContext instance for the Perl runtime.
		virtual std::shared_ptr<IRuntimeContext> Perl(std::string configName = "default") = 0;

		/// \brief Creates IRuntimeContext instance to interact with Python runtime.
		/// \param configName The name of the configuration to use.
		/// \return A shared_ptr to an IRuntimeContext instance for the Python runtime.
		virtual std::shared_ptr<IRuntimeContext> Python(std::string configName = "default") = 0;

		/// \brief Creates IRuntimeContext instance to interact with Ruby runtime.
		/// \param configName The name of the configuration to use.
		/// \return A shared_ptr to an IRuntimeContext instance for the Ruby runtime.
		virtual std::shared_ptr<IRuntimeContext> Ruby(std::string configName = "default") = 0;

		/// \brief Creates IRuntimeContext instance to interact with Node.js runtime.
		/// \param configName The name of the configuration to use.
		/// \return A shared_ptr to an IRuntimeContext instance for the Node.js runtime.
		virtual std::shared_ptr<IRuntimeContext> Nodejs(std::string configName = "default") = 0;
	};
}