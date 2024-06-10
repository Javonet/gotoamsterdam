#pragma once
#include "Interfaces/IRuntimeContext.h"

namespace JavonetNS::Cpp::Sdk::Interfaces {
	/// \brief The IRuntimeFactory interface provides methods for creating runtime contexts.
	/// Each method corresponds to a specific runtime (CLR, JVM, .NET Core, Perl, Ruby, Node.js, Python) and returns a shared_ptr to an IRuntimeContext instance for that runtime.
	class IRuntimeFactory {
	public:
		virtual ~IRuntimeFactory() = default;

		/// \brief Creates IRuntimeContext instance to interact with CLR runtime.
		/// \return A shared_ptr to an IRuntimeContext instance for the CLR runtime.
		virtual std::shared_ptr<IRuntimeContext> Clr() = 0;

		/// \brief Creates IRuntimeContext instance to interact with JVM runtime.
		/// \return A shared_ptr to an IRuntimeContext instance for the JVM runtime.
		virtual std::shared_ptr<IRuntimeContext> Jvm() = 0;

		/// \brief Creates IRuntimeContext instance to interact with .NET Core runtime.
		/// \return A shared_ptr to an IRuntimeContext instance for the .NET Core runtime.
		virtual std::shared_ptr<IRuntimeContext> Netcore() = 0;

		/// \brief Creates IRuntimeContext instance to interact with Perl runtime.
		/// \return A shared_ptr to an IRuntimeContext instance for the Perl runtime.
		virtual std::shared_ptr<IRuntimeContext> Perl() = 0;

		/// \brief Creates IRuntimeContext instance to interact with Python runtime.
		/// \return A shared_ptr to an IRuntimeContext instance for the Python runtime.
		virtual std::shared_ptr<IRuntimeContext> Python() = 0;

		/// \brief Creates IRuntimeContext instance to interact with Ruby runtime.
		/// \return A shared_ptr to an IRuntimeContext instance for the Ruby runtime.
		virtual std::shared_ptr<IRuntimeContext> Ruby() = 0;

		/// \brief Creates IRuntimeContext instance to interact with Node.js runtime.
		/// \return A shared_ptr to an IRuntimeContext instance for the Node.js runtime.
		virtual std::shared_ptr<IRuntimeContext> Nodejs() = 0;
	};
}