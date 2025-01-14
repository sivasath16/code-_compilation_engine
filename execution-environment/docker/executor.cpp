#include <cstdlib>
#include <fstream>
#include <iostream>
#include <sstream>
#include <string>
#include <stdexcept>

int main() {
    try {
        // Get the user-provided code from the CODE environment variable
        const char* code = std::getenv("CODE");
        if (!code) {
            std::cerr << "Error: No code provided." << std::endl;
            return 1;
        }

        // Write the code to a file
        std::ofstream file("program.cpp");
        if (!file.is_open()) {
            std::cerr << "Error: Unable to write to file." << std::endl;
            return 1;
        }
        file << code;
        file.close();

        // Compile the C++ file
        std::cout << "Compiling the program..." << std::endl;
        int compileStatus = std::system("g++ program.cpp -o program 2> compile_errors.txt");
        if (compileStatus != 0) {
            // Read and output the compilation errors
            std::ifstream compileErrors("compile_errors.txt");
            std::stringstream errorStream;
            errorStream << compileErrors.rdbuf();
            compileErrors.close();
            std::cerr << "Error: Compilation failed.\n" << errorStream.str() << std::endl;
            return 1;
        }

        // Run the compiled program
        std::cout << "Running the program..." << std::endl;
        int runStatus = std::system("./program > program_output.txt 2> runtime_errors.txt");
        if (runStatus != 0) {
            // Read and output the runtime errors
            std::ifstream runtimeErrors("runtime_errors.txt");
            std::stringstream errorStream;
            errorStream << runtimeErrors.rdbuf();
            runtimeErrors.close();
            std::cerr << "Error: Execution failed.\n" << errorStream.str() << std::endl;
            return 1;
        }

        // Read and output the program's successful output
        std::ifstream programOutput("program_output.txt");
        std::stringstream outputStream;
        outputStream << programOutput.rdbuf();
        programOutput.close();
        std::cout << outputStream.str() << std::endl;

    } catch (const std::exception& e) {
        std::cerr << "Error: " << e.what() << std::endl;
        return 1;
    } catch (...) {
        std::cerr << "Error: Unknown exception occurred." << std::endl;
        return 1;
    }

    return 0;
}
