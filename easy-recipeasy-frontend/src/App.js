import React, { useEffect, useState } from "react";
import axios from "axios";
import { BrowserRouter as Router, Route, Routes } from "react-router-dom";
import Dashboard from "./components/Dashboard";

const App = () => {
  const [healthData, setHealthData] = useState(null);

  const fetchHealth = async () => {
    try {
      const response = await axios.get("http://localhost:8080/health"); // Update URL
      setHealthData(response.data);
    } catch (error) {
      console.error("Error fetching health check:", error);
    }
  };

  const handleLogin = () => {
    window.location.href = "http://localhost:8080/auth/google/login";
  };

  useEffect(() => {
    fetchHealth();
  }, []);

  return (
    <Router>
      <Routes>
        <Route path="/dashboard" element={<Dashboard />} />
        <Route path="/" element={
          <div className="flex flex-col justify-center items-center h-screen bg-gray-100">
            <h1 className="font-bold text-4xl">Welcome to Easy Recipeasy!</h1>
            {healthData ? (
              <div className="mt-4">
                <p className="text-lg">Server Status: {healthData.status}</p>
              </div>
            ) : (
              <p className="mt-4 text-lg">Loading... (Backend Down)</p>
            )}
            <button
              onClick={handleLogin}
              className="mt-4 px-4 py-2 bg-blue-600 text-white rounded"
            >
              Sign in with Google
            </button>
          </div>
        } />
      </Routes>
    </Router>
  );
};

export default App;