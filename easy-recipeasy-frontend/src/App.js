import React, { useEffect, useState } from "react";
import axios from "axios";

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

  useEffect(() => {
    fetchHealth();
  }, []);

  return (
    <div className="flex flex-col justify-center items-center h-screen bg-gray-100">
      <h1 className="font-bold text-4xl">Welcome to Easy Recipeasy!</h1>
      {healthData ? (
        <div className="mt-4">
          <p className="text-lg">Server Status: {healthData.status}</p>
        </div>
      ) : (
        <p className="mt-4 text-lg">Loading... (Backend Down)</p>
      )}
    </div>
  );
};

export default App;