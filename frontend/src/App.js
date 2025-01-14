import React, { useState } from "react";
import Editor from "@monaco-editor/react";
import axios from "axios";
import "./App.css";

const App = () => {
  const [code, setCode] = useState("# Write your code here...");
  const [language, setLanguage] = useState("python");
  const [output, setOutput] = useState("");
  const [loading, setLoading] = useState(false);

  const languages = [
    { label: "JavaScript", value: "javascript" },
    { label: "Python", value: "python" },
    { label: "Java", value: "java" },
    { label: "C++", value: "cpp" },
  ];

  const languageTemplates = {
    java: `class Main {
    public static void main(String[] args) {
        // Write your code here!!
    }
}`,
    python: `# Write your code here!!`,
    javascript: `// Write your code here!!`,
    cpp: `#include <iostream>
using namespace std;

int main() {
    // Write your code here!!
    return 0;
}`,
  };

  const handleLanguageChange = (e) => {
    const selectedLanguage = e.target.value;
    setLanguage(selectedLanguage);

    // Set the template code for the selected language
    setCode(languageTemplates[selectedLanguage] || "// Unsupported language");
  };

  const handleExecute = async () => {
    setLoading(true);
    setOutput("Executing...");
    try {
      // Send the code and language to the backend
      const response = await axios.post("http://localhost:8080/api/execute", {
        code,
        language,
      });
  
      console.log("Execution response:", response);
  
      // Handle the response:
      if (response.data.output) {
        // If the backend provides immediate output
        setOutput(response.data.output || "Execution completed!");
      } else if (response.data.error) {
        // If there's an error in the response
        setOutput(`Error: ${response.data.error}`);
      } else if (response.data.task_id) {
        // If the backend provides a task_id for asynchronous processing
        const taskId = response.data.task_id;
  
        // Poll for results every 2 seconds
        const interval = setInterval(async () => {
          try {
            const resultResponse = await axios.get(
              `http://localhost:8080/api/results/${taskId}`
            );
  
            console.log("Polling result response:", resultResponse);
  
            if (resultResponse.data.output || resultResponse.data.error) {
              // Display the output or error from the result
              setOutput(
                resultResponse.data.output ||
                  `Error: ${resultResponse.data.error}`
              );
              clearInterval(interval); // Stop polling once we get a result
            }
          } catch (pollError) {
            console.error("Error polling for result:", pollError);
            setOutput("Error fetching execution result.");
            clearInterval(interval); // Stop polling on error
          }
        }, 2000); // Poll every 2 seconds
      } else {
        setOutput("Unknown response from the server.");
      }
    } catch (error) {
      console.error("Execution error:", error);
      // Display errors from the backend or generic error message
      setOutput(
        error.response?.data?.error || error.response?.data?.message || "Error executing code."
      );
    } finally {
      setLoading(false); // Ensure loading state is cleared
    }
  };
  


  

  return (
    <div className="app-container">
      <h1>Code Execution Engine</h1>
  
      {/* Language Selector */}
      <div className="language-selector">
        <label htmlFor="language">Select Language:</label>
        <select id="language" value={language} onChange={handleLanguageChange}>
          {languages.map((lang) => (
            <option key={lang.value} value={lang.value}>
              {lang.label}
            </option>
          ))}
        </select>
      </div>
  
      {/* Editor and Output Side-by-Side */}
      <div className="editor-output-container">
        {/* Code Editor */}
        <div className="editor-container">
          <Editor
            height="400px"
            language={language}
            theme="vs-dark"
            value={code}
            onChange={(newCode) => setCode(newCode || "")}
          />
        </div>
  
        {/* Output Section */}
        <div className={`output-container ${
            output.startsWith("Error:") ? "error" : ""
          }`}>
          <h2>Output:</h2>
          <pre>{output}</pre>
        </div>
      </div>
  
      {/* Run Code Button */}
      <div className="button-container">
        <button
          className="execute-btn"
          onClick={handleExecute}
          disabled={loading}
        >
          {loading ? "Executing..." : "Run Code"}
        </button>
      </div>
    </div>
  );
  
};

export default App;