// src/components/GoogleLoginButton.js
import React from 'react';
import { GoogleLogin } from '@react-oauth/google';
import { jwtDecode } from 'jwt-decode';
import axios from 'axios';

const GoogleLoginButton = () => {
  const handleLoginSuccess = async (credentialResponse) => {
    if (credentialResponse.credential) {
      const decoded = jwtDecode(credentialResponse.credential);
      console.log('User Information:', decoded);
  
      try {
        const res = await axios.post('/api/auth/google', {
          token: credentialResponse.credential,
        });
        // Handle response (e.g., save tokens, redirect)
        console.log('Server Response:', res.data);
      } catch (error) {
        console.error('Error sending token to backend:', error);
      }
    }
  };

  const handleLoginError = () => {
    console.error('Login Failed');
  };

  return (
    <div className="flex justify-center mt-4">
      <GoogleLogin
        onSuccess={handleLoginSuccess}
        onError={handleLoginError}
        useOneTap
      />
    </div>
  );
};

export default GoogleLoginButton;