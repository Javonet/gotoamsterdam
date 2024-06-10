#pragma once
#include "commonHeaders.h"

namespace JavonetNS::Cpp::Sdk::Interfaces {
	class IInvocationContext : public std::enable_shared_from_this<IInvocationContext> {
	public:
		virtual ~IInvocationContext() {};
		 /// \brief Executes the current command.
		 /// The initial state of the invocation context, which we call non-materialized, wraps either a single command or a chain of recursively nested commands.
		 /// Commands become nested through each invocation of methods on the Invocation Context. Each invocation triggers the creation of a new Invocation Context instance that wraps the current command with a new parent command valid for the invoked method.
		 /// The developer can decide at any moment to materialize the context, taking full control of the chunks of the expression being transferred and processed on the target runtime.
		 /// Refer to this <a href="https://www.javonet.com/guides/v2/cpp/foundations/execute-method">article on Javonet Guides</a>
		 /// \return The InvocationContext after executing the command.
		virtual std::shared_ptr<IInvocationContext> Execute() = 0;

		/// \brief Invokes a static method on the target runtime.
		/// This method creates a new InvocationContext instance that wraps the command to invoke the static method.
		/// Refer to this <a href="https://www.javonet.com/guides/v2/cpp/calling-methods/invoking-static-method">article on Javonet Guides</a>
		/// \param methodName The name of the method to invoke.
		/// \param arguments The arguments to pass to the static method.
		/// \return A new InvocationContext instance that wraps the command to invoke the static method.
		virtual std::shared_ptr<IInvocationContext> InvokeStaticMethod(std::string methodName, std::deque<std::any> arguments) = 0;
		/// \brief Invokes a static method on the target runtime.
		/// This method creates a new InvocationContext instance that wraps the command to invoke the static method.
		/// Refer to this <a href="https://www.javonet.com/guides/v2/cpp/calling-methods/invoking-static-method">article on Javonet Guides</a>
		/// \param methodName The name of the method to invoke.
		/// \param argument The argument to pass to the static method.
		/// \return A new InvocationContext instance that wraps the command to invoke the static method.
		virtual std::shared_ptr<IInvocationContext> InvokeStaticMethod(std::string methodName, std::any argument) = 0;
		virtual std::shared_ptr<IInvocationContext> InvokeStaticMethod(std::string methodName) = 0;
        /// \brief Invokes an instance method on the target runtime.
		 /// This method creates a new InvocationContext instance that wraps the command to invoke the instance method.
		/// Refer to this <a href="https://www.javonet.com/guides/v2/cpp/calling-methods/creating-instance-and-calling-instance-methods">article on Javonet Guides</a>
        /// \param methodName The name of the method to invoke.
        /// \param arguments The arguments to pass to the instance method.
        /// \return A new InvocationContext instance that wraps the command to invoke the instance method.
		virtual std::shared_ptr<IInvocationContext> InvokeInstanceMethod(std::string methodName, std::deque<std::any> arguments) = 0;
		/// \brief Invokes an instance method on the target runtime.
		/// This method creates a new InvocationContext instance that wraps the command to invoke the instance method.
		/// Refer to this <a href="https://www.javonet.com/guides/v2/cpp/calling-methods/creating-instance-and-calling-instance-methods">article on Javonet Guides</a>
		/// \param methodName The name of the method to invoke.
		/// \param argument The argument to pass to the instance method.
		/// \return A new InvocationContext instance that wraps the command to invoke the instance method.
		virtual std::shared_ptr<IInvocationContext> InvokeInstanceMethod(std::string methodName, std::any argument) = 0;
		virtual std::shared_ptr<IInvocationContext> InvokeInstanceMethod(std::string methodName) = 0;
		/// \brief Retrieves the value of a static field from the target runtime.
		/// This method creates a new InvocationContext instance that wraps the command to get the static field.
		/// Refer to this <a href="https://www.javonet.com/guides/v2/cpp/fields-and-properties/getting-and-setting-values-for-static-fields-and-properties">article on Javonet Guides</a>
		/// \param fieldName The name of the field to retrieve.
		/// \return A new InvocationContext instance that wraps the command to get the static field.
        virtual std::shared_ptr<IInvocationContext> GetStaticField(std::string fieldName) = 0;
		/// \brief Sets the value of a static field in the target runtime.
		/// This method creates a new InvocationContext instance that wraps the command to set the static field.
		/// Refer to this <a href="https://www.javonet.com/guides/v2/cpp/fields-and-properties/getting-and-setting-values-for-static-fields-and-properties">article on Javonet Guides</a>
		/// \param fieldName The name of the field to set.
		/// \param value The new value to set.
		/// \return A new InvocationContext instance that wraps the command to set the static field.
		virtual std::shared_ptr<IInvocationContext> SetStaticField(std::string fieldName, std::any value) = 0;
		/// \brief Creates a new instance of a class in the target runtime.
		/// This method creates a new InvocationContext instance that wraps the command to create the instance.
		/// Refer to this <a href="https://www.javonet.com/guides/v2/cpp/calling-methods/creating-instance-and-calling-instance-methods">article on Javonet Guides</a>
		/// \param arguments The class constructor arguments.
		/// \return A new InvocationContext instance that wraps the command to create the instance.
		virtual std::shared_ptr<IInvocationContext> CreateInstance(std::deque<std::any> arguments) = 0;
		/// \brief Creates a new instance of a class in the target runtime.
		/// This method creates a new InvocationContext instance that wraps the command to create the instance.
		/// Refer to this <a href="https://www.javonet.com/guides/v2/cpp/calling-methods/creating-instance-and-calling-instance-methods">article on Javonet Guides</a>
		/// \param argument The class constructor argument.
		/// \return A new InvocationContext instance that wraps the command to create the instance.
		virtual std::shared_ptr<IInvocationContext> CreateInstance(std::any argument) = 0;
		virtual std::shared_ptr<IInvocationContext> CreateInstance() = 0;
		/// \brief Gets the value of instance field from the target runtime.
		/// This method creates a new InvocationContext instance that wraps the command to get the instance field.
		/// Refer to this <a href="https://www.javonet.com/guides/v2/fields-and-properties/getting-and-setting-values-for-instance-fields-and-properties">article on Javonet Guides</a>
		/// \param fieldName The name of the property to retrieve.
		/// \return A new InvocationContext instance that wraps the command to get the instance field.
		virtual std::shared_ptr<IInvocationContext> GetInstanceField(std::string fieldName) = 0;
		/// \brief Sets the value of an instance field in the target runtime.
		/// This method creates a new InvocationContext instance that wraps the command to set the instance field.
		/// Refer to this <a href="https://www.javonet.com/guides/v2/fields-and-properties/getting-and-setting-values-for-instance-fields-and-properties">article on Javonet Guides</a>
		/// \param fieldName The name of the field to set.
		/// \param value The new value to set.
		/// \return A new InvocationContext instance that wraps the command to set the instance field.
		virtual std::shared_ptr<IInvocationContext> SetInstanceField(std::string fieldName, std::any value) = 0;
		/// \brief Gets the element from an array or a collection in the target runtime.
		/// This method creates a new InvocationContext instance that wraps the command to get the index.
		/// Refer to this <a href="https://www.javonet.com/guides/v2/cpp/arrays-and-collections/one-dimensional-arrays">article on Javonet Guides</a>
		/// \param indexes The indexes of the element to retrieve.
		/// \return A new InvocationContext instance that wraps the command to get the index.
		virtual std::shared_ptr<IInvocationContext> GetIndex(std::vector<std::any> indexes) = 0;
		/// \brief Gets the element from an array or a collection in the target runtime.
		/// This method creates a new InvocationContext instance that wraps the command to get the index.
		/// Refer to this <a href="https://www.javonet.com/guides/v2/cpp/arrays-and-collections/one-dimensional-arrays">article on Javonet Guides</a>
		/// \param index The index of the element to retrieve.
		/// \return A new InvocationContext instance that wraps the command to get the index.
		virtual std::shared_ptr<IInvocationContext> GetIndex(std::any index) = 0;
		/// \brief Sets the element in an array or a collection in the target runtime.
		/// \param indexes The indexes of the element to set.
		/// \param value The new value to set.
		/// This method creates a new InvocationContext instance that wraps the command to set the index.
		/// Refer to this <a href="https://www.javonet.com/guides/v2/cpp/arrays-and-collections/one-dimensional-arrays">article on Javonet Guides</a>
		/// \return A new InvocationContext instance that wraps the command to set the index.
		virtual std::shared_ptr<IInvocationContext> SetIndex(std::vector<std::any> indexes, std::any value) = 0;
		/// \brief Sets the element in an array or a collection in the target runtime.
		/// \param index The index of the element to set.
		/// \param value The new value to set.
		/// This method creates a new InvocationContext instance that wraps the command to set the index.
		/// Refer to this <a href="https://www.javonet.com/guides/v2/cpp/arrays-and-collections/one-dimensional-arrays">article on Javonet Guides</a>
		/// \return A new InvocationContext instance that wraps the command to set the index.
		virtual std::shared_ptr<IInvocationContext> SetIndex(std::any index, std::any value) = 0;
		/// \brief Get number of elements of an array or collection in the target runtime.
		/// This method creates a new InvocationContext instance that wraps the command to get the size.
		/// Refer to this <a href="https://www.javonet.com/guides/v2/cpp/arrays-and-collections/arrays">article on Javonet Guides</a>
		/// \return A new InvocationContext instance that wraps the command to get the size.
		virtual std::shared_ptr<IInvocationContext> GetSize() = 0;
		/// \brief Get the rank of an array in the target runtime.
		/// This method creates a new InvocationContext instance that wraps the command to get the rank.
		/// Refer to this <a href="https://www.javonet.com/guides/v2/cpp/arrays-and-collections/arrays">article on Javonet Guides</a>
		/// \return A new InvocationContext instance that wraps the command to get the rank.
		virtual std::shared_ptr<IInvocationContext> GetRank() = 0;
		/// \brief Invokes a generic static method on the target runtime.
		/// This method creates a new InvocationContext instance that wraps the command to invoke the generic static method.
		/// Refer to this <a href="https://www.javonet.com/guides/v2/cpp/generics/calling-generic-static-method">article on Javonet Guides</a>
		/// \param methodName The name of the method to invoke.
		/// \param arguments The arguments to pass to the invoke generic static method command.
		virtual std::shared_ptr<IInvocationContext> InvokeGenericStaticMethod(std::string methodName, std::deque<std::any> arguments) = 0;
		/// \brief Invokes a generic static method on the target runtime.
		/// This method creates a new InvocationContext instance that wraps the command to invoke the generic static method.
		/// Refer to this <a href="https://www.javonet.com/guides/v2/cpp/generics/calling-generic-static-method">article on Javonet Guides</a>
		/// \param methodName The name of the method to invoke.
		/// \param argument The argument to pass to the invoke generic static method command.
		virtual std::shared_ptr<IInvocationContext> InvokeGenericStaticMethod(std::string methodName, std::any argument) = 0;
		virtual std::shared_ptr<IInvocationContext> InvokeGenericStaticMethod(std::string methodName) = 0;
		/// \brief Invokes a generic instance method on the target runtime.
		/// \param methodName The name of the method to invoke.
		/// \param arguments The arguments to pass to the invoke generic instance method command.
		/// This method creates a new InvocationContext instance that wraps the command to invoke the generic instance method.
		/// Refer to this <a href="https://www.javonet.com/guides/v2/cpp/generics/calling-generic-instance-method">article on Javonet Guides</a>
		/// \return A new InvocationContext instance that wraps the command to invoke the generic instance method.
		virtual std::shared_ptr<IInvocationContext> InvokeGenericMethod(std::string methodName, std::deque<std::any> arguments) = 0;
		/// \brief Invokes a generic instance method on the target runtime.
		/// \param methodName The name of the method to invoke.
		/// \param argument The argument to pass to the invoke generic instance method command.
		/// This method creates a new InvocationContext instance that wraps the command to invoke the generic instance method.
		/// Refer to this <a href="https://www.javonet.com/guides/v2/cpp/generics/calling-generic-instance-method">article on Javonet Guides</a>
		/// \return A new InvocationContext instance that wraps the command to invoke the generic instance method.
		virtual std::shared_ptr<IInvocationContext> InvokeGenericMethod(std::string methodName, std::any argument) = 0;
		/// \brief Invokes a generic instance method on the target runtime.
		/// \param methodName The name of the method to invoke.
		/// This method creates a new InvocationContext instance that wraps the command to invoke the generic instance method.
		/// Refer to this <a href="https://www.javonet.com/guides/v2/cpp/generics/calling-generic-instance-method">article on Javonet Guides</a>
		/// \return A new InvocationContext instance that wraps the command to invoke the generic instance method.
		virtual std::shared_ptr<IInvocationContext> InvokeGenericMethod(std::string methodName) = 0;
		/// \brief Retrieves the name of an enum from the target runtime.
		/// This method creates a new InvocationContext instance that wraps the command to get the enum name.
		/// Refer to this <a href="https://www.javonet.com/guides/v2/cpp/enums/using-enum-type">article on Javonet Guides</a>
		/// \return A new InvocationContext instance that wraps the command to get the enum name.
		/// See also: \ref [article on Javonet Guides](https://www.javonet.com/guides/v2/cpp/enums/using-enum-type)
		virtual std::shared_ptr<IInvocationContext> GetEnumName() = 0;
		/// \brief Retrieves the value of an enum from the target runtime.
		/// This method creates a new InvocationContext instance that wraps the command to get the enum value.
		/// Refer to this <a href="https://www.javonet.com/guides/v2/cpp/enums/using-enum-type">article on Javonet Guides</a>
		/// \return A new InvocationContext instance that wraps the command to get the enum value.
		virtual std::shared_ptr<IInvocationContext> GetEnumValue() = 0;
		/// \brief Retrieves the value of a reference type from the target runtime.
		/// This method creates a new InvocationContext instance that wraps the command to get the reference value.
		/// Refer to this <a href="https://www.javonet.com/guides/v2/cpp/methods-arguments/passing-arguments-by-reference-with-ref-keyword">article on Javonet Guides</a>
		/// \return A new InvocationContext instance that wraps the command to get the reference value.
		virtual std::shared_ptr<IInvocationContext> GetRefValue() = 0;
		/// \brief Retrieves an array from the target runtime.
		/// This method creates a new InvocationContext instance that wraps the command to get the array.
		/// Refer to this <a href="https://www.javonet.com/guides/v2/cpp/arrays-and-collections/one-dimensional-arrays">article on Javonet Guides</a>
		/// \return A new InvocationContext instance that wraps the command to get the array.
		virtual std::any RetrieveArray() = 0;
		/// \brief Retrieves the value from a command
		/// \return The value of the command.
		virtual std::any GetValue() = 0;
		virtual std::shared_ptr<IInvocationContext> operator[](std::any index) = 0;
		virtual std::shared_ptr<IInvocationContext> operator[](std::vector<std::any> indexes) = 0;
	};
}