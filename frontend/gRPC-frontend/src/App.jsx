import React, { useEffect, useState } from 'react';
import axios from 'axios';
import ChatBox from './components/Chatbox';
import Input from './components/Input';


const App = () => {
  const [messages, setMessages] = useState([]);

  const fetchMessages = async () => {
    try {
      const res = await axios.get('http://localhost:8080/v1/chat/messages');
      setMessages(res.data.messages || []);
    } catch (error) {
      console.error('Error fetching messages:', error);
    }
  };

  //poll every 2 seconds
  useEffect(() => {
    fetchMessages(); 
    const interval = setInterval(fetchMessages, 2000);
    return () => clearInterval(interval);
  }, []);

  return (
    <div className="App">
      <ChatBox messages={messages} />
          <Input onSend={fetchMessages} />

    </div>
  );
};

export default App;
