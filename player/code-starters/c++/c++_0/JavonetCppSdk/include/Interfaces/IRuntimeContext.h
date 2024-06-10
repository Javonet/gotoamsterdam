#pragma once 
#include "IInvocationContext.h"
#undef LoadLibrary

namespace JavonetNS::Cpp::Sdk::Interfaces {
	/// \brief Represents a single context which allows interaction with a selected technology.
	/// This class refers to a single instance of the called runtime within a particular target OS process. This can be either the local currently running process (inMemory) or a particular remote process identified by the IP Address and PORT of the target Javonet instance.
	/// Multiple Runtime Contexts can be initialized within one process. Calling the same technology on inMemory communication channel will return the existing instance of runtime context. Calling the same technology on TCP channel but on different nodes will result in unique Runtime Contexts.
	/// Within the runtime context, any number of libraries can be loaded and any objects from the target technology can be interacted with, as they are aware of each other due to sharing the same memory space and same runtime instance.
	/// Refer to this <a href="https://www.javonet.com/guides/v2/cpp/foundations/runtime-context">article on Javonet Guides</a>
	class IRuntimeContext : public std::enable_shared_from_this<IRuntimeContext> {
	public:
		virtual ~IRuntimeContext() {};

		/// \brief Adds a reference to a library. Javonet allows you to reference and use modules or packages written in various languages.
		/// This method allows you to use any library from all supported technologies. The necessary libraries need to be referenced.
		/// The argument is a relative or full path to the library. If the library has dependencies on other libraries, the latter needs to be added first.
		/// After referencing the library, any objects stored in this package can be used. Use static classes, create instances, call methods, use fields and properties, and much more.
		/// Refer to this <a href="https://www.javonet.com/guides/v2/cpp/getting-started/adding-references-to-libraries">article on Javonet Guides</a>
		/// \param libraryPath The relative or full path to the library.
		/// \return RuntimeContext instance.
		virtual std::shared_ptr<IRuntimeContext> LoadLibrary(std::string libraryPath) = 0;

		/// \brief Retrieves a reference to a specific type. 
		/// The type can be a class, interface or enum. The type can be retrieved from any referenced library.
		/// \param typeName The full name of the type.
		/// \param arguments The arguments to be passed, if needed
		/// \return InvocationContext instance that wraps the command to get type.
		virtual std::shared_ptr<IInvocationContext> GetType(std::string typeName, std::deque<std::any> arguments) = 0;

		/// \brief Retrieves a reference to a specific type.
		/// The type can be a class, interface or enum. The type can be retrieved from any referenced library.
		/// \param typeName The full name of the type.
		/// \param argument The argument to be passed, if needed
		/// \return InvocationContext instance that wraps the command to get type.
		virtual std::shared_ptr<IInvocationContext> GetType(std::string typeName, std::any argument) = 0;
		virtual std::shared_ptr<IInvocationContext> GetType(std::string typeName) = 0;

		/// \brief Casts the provided value to a specific type. This method is used when invoking methods that require specific types of arguments.
		/// The arguments include the target type and the value to be cast. The target type must be retrieved from the called runtime using the GetType method.
		/// After casting the value, it can be used as an argument when invoking methods.
		/// Refer to this <a href="https://www.javonet.com/guides/v2/cpp/casting/casting">article on Javonet Guides</a>
		/// \param arguments The target type and the value to be cast.
		/// \return InvocationContext instance that wraps the command to cast argument.
		virtual std::shared_ptr<IInvocationContext> Cast(std::deque<std::any> arguments) = 0;

		/// \brief Retrieves a specific item from an enum type. This method is used when working with enums from the called runtime.
		/// The arguments include the enum type and the name of the item. The enum type must be retrieved from the called runtime using the GetType method.
		/// After retrieving the item, it can be used as an argument when invoking methods or for other operations.
		/// Refer to this <a href="https://www.javonet.com/guides/v2/cpp/enums/using-enum-type">article on Javonet Guides</a>
		/// \param arguments The enum type and the name of the item.
		/// \return InvocationContext instance that wraps command to get enum item.
		virtual std::shared_ptr<IInvocationContext> GetEnumItem(std::deque<std::any> arguments) = 0;
		virtual std::shared_ptr<IInvocationContext> GetEnumItem(std::any argument) = 0;

		/// \brief Creates a reference type argument that can be passed to a method with a ref parameter modifier. This method is used when working with methods from the called runtime that require arguments to be passed by reference.
		/// The arguments include the value and optionally the type of the reference. The type must be retrieved from the called runtime using the GetType method.
		/// After creating the reference, it can be used as an argument when invoking methods.
		/// Refer to this <a href="https://www.javonet.com/guides/v2/cpp/methods-arguments/passing-arguments-by-reference-with-ref-keyword">article on Javonet Guides</a>
		/// \param arguments The value and optionally the type of the reference.
		/// \return InvocationContext instance that wraps the command to create reference as ref parameter.
		virtual std::shared_ptr<IInvocationContext> AsRef(std::deque<std::any> arguments) = 0;
		virtual std::shared_ptr<IInvocationContext> AsRef(std::any argument) = 0;

		/// \brief Creates a reference type argument that can be passed to a method with an out parameter modifier. This method is used when working with methods from the called runtime that require arguments to be passed by reference.
		/// The arguments include the value and optionally the type of the reference. The type must be retrieved from the called runtime using the GetType method.
		/// After creating the reference, it can be used as an argument when invoking methods.
		/// \param arguments The value and optionally the type of the reference.
		/// \return InvocationContext instance that wraps the command to create reference as out parameter.
		/// \see <a href="https://www.javonet.com/guides/v2/cpp/methods-arguments/passing-arguments-by-reference-with-out-keyword">Passing Arguments by Reference with out Keyword Guide</a>
		virtual std::shared_ptr<IInvocationContext> AsOut(std::deque<std::any> arguments) = 0;
		virtual std::shared_ptr<IInvocationContext> AsOut(std::any argument) = 0;
	private:
		virtual void Execute() = 0;
	};
}