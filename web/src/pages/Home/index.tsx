import React, { useState } from 'react';
import { Box } from '@mui/material';
import { grey, red } from '@mui/material/colors';
import { MessageCard, MessageCardProps } from '../../components/Card/index';

const Home: React.FC = () => {
  const [messages, setMessages] = useState<Array<MessageCardProps>>([]);
  const socket = new WebSocket('ws://127.0.0.1:8000/chat');
  socket.onopen = (event) => {
    alert('conection');
  };
  socket.addEventListener('message', (event) => {
    console.log(event.data);
    const payload = JSON.parse(event.data);
    const newMessage: MessageCardProps = {
      id: payload.id,
      userID: payload.user_id,
      text: payload.text,
      createdAt: payload.created_at,
    };
    setMessages((prevMessages: MessageCardProps[]) => [
      ...prevMessages,
      newMessage,
    ]);
  });

  return (
    <Box
      sx={{
        marginTop: '60px',
      }}
    >
      {messages.map((message, index) => (
        <MessageCard
          id={message.id}
          userID={message.userID}
          text={message.text}
          createdAt={message.createdAt}
        />
      ))}
    </Box>
  );
};

export default Home;
