import React, { useState } from 'react';
import { Box } from '@mui/material';
import { MessageCard, MessageCardProps } from '../../components/Card/index';
import InputForm from '../../components/InputForm/index';

const Home: React.FC = () => {
  const [messages, setMessages] = useState<Array<MessageCardProps>>([]);
  const socket = new WebSocket('ws://127.0.0.1:8000/chat');

  socket.addEventListener('message', (event) => {
    console.log(event.data);
    const payload = JSON.parse(event.data);
    const date = new Date(payload.created_at);
    const datetime = `${date.getFullYear()}/${date.getMonth()}/${date.getDate()} ${date.getHours()}:${date.getMinutes()}`;
    const newMessage: MessageCardProps = {
      id: payload.id,
      userID: payload.user_id,
      text: payload.text,
      createdAt: datetime,
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
      <InputForm />
    </Box>
  );
};

export default Home;
