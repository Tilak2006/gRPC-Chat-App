import React from 'react';
import './ChatBox.css';

const ChatBox = ({ messages }) => {
  return (
    <div className="chat-box">
      <h2>Chat Messages</h2>
      <div className="message-list">
        {messages.length === 0 ? (
          <p className="empty">No messages yet</p>
        ) : (
          messages.map((msg, index) => (
            <div className="message" key={index}>
              <strong>{msg.senderName}</strong>: {msg.message}
              <div className="timestamp">{new Date(msg.timestamp * 1000).toLocaleTimeString()}</div>
            </div>
          ))
        )}
      </div>
    </div>
  );
};

export default ChatBox;
