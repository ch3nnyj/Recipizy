import React, { useEffect } from "react";

const Dashboard = () => {
  useEffect(() => {
    const params = new URLSearchParams(window.location.search);
    const token = params.get("token");
    if (token) {
      localStorage.setItem("authToken", token);
    }
    // Redirect to clean up URL
    window.history.replaceState({}, document.title, "/dashboard");
  }, []);

  return (
    <div>
      <h2>Welcome to your dashboard!</h2>
      {/* Render dashboard content */}
    </div>
  );
};

export default Dashboard;