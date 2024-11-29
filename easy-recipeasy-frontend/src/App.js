import React, { useEffect } from "react";
import axios from "axios";
import GoogleLoginButton from "./components/GoogleLoginButton";

const App = () => {
  const fetchHealth = async () => {
    try {
      const response = await axios.get("/api/health");
      console.log(response.data);
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
      <GoogleLoginButton />
    </div>
  );
};

export default App;