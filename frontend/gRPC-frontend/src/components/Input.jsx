import React, { useState } from 'react';
import axios from 'axios';
import './Input.css';

const Input = ({ onSend }) => {
  const [message, setMessage] = useState('');

  const handleSend = async () => {
    if (!message.trim()) return;

    try {
      await axios.post('http://localhost:8080/v1/chat/send', {
        sender_id: 'u1', //temporary hardcoded, will make dynamic later
        sender_name: 'Tilak',
        message: message.trim(),
      });

      setMessage('');
      if (onSend) onSend(); 
    } catch (err) {
      console.error('Error sending message:', err);
    }
  };

  return (
    <div className="input-bar">
      <input
        type="text"
        placeholder="Type your message..."
        value={message}
        onChange={(e) => setMessage(e.target.value)}
        onKeyDown={(e) => e.key === 'Enter' && handleSend()}
      />
      <button onClick={handleSend}>Send</button>
    </div>
  );
};

export default Input;
