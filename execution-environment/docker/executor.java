import java.io.*;
import java.nio.file.*;

public class executor {
    public static void main(String[] args) {
        try {
            // Get the code from the CODE environment variable
            String code = System.getenv("CODE");

            // Write the code to a file
            Files.write(Paths.get("Main.java"), code.getBytes());

            // Compile the Java file
            Process compileProcess = Runtime.getRuntime().exec("javac Main.java");
            compileProcess.waitFor();

            if (compileProcess.exitValue() != 0) {
                // Capture compilation errors
                BufferedReader errorReader = new BufferedReader(new InputStreamReader(compileProcess.getErrorStream()));
                StringBuilder errors = new StringBuilder();
                String line;
                while ((line = errorReader.readLine()) != null) {
                    errors.append(line).append("\n");
                }
                System.err.println(errors.toString()); // Output to stderr
                return;
            }

            // Run the compiled Java program
            Process runProcess = Runtime.getRuntime().exec("java Main");
            BufferedReader outputReader = new BufferedReader(new InputStreamReader(runProcess.getInputStream()));
            BufferedReader errorReader = new BufferedReader(new InputStreamReader(runProcess.getErrorStream()));

            StringBuilder output = new StringBuilder();
            String line;

            // Capture runtime errors (stderr)
            StringBuilder runtimeErrors = new StringBuilder();
            while ((line = errorReader.readLine()) != null) {
                runtimeErrors.append(line).append("\n");
            }

            // Capture successful output (stdout)
            while ((line = outputReader.readLine()) != null) {
                output.append(line).append("\n");
            }

            // Print output or errors
            if (runtimeErrors.length() > 0) {
                System.err.println(runtimeErrors.toString());
            } else {
                System.out.println(output.toString());
            }
        } catch (Exception e) {
            System.err.println("Error: " + e.getMessage());
        }
    }
}
