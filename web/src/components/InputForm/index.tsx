import React, { ChangeEvent, useState } from 'react';
import { Box, Button, TextField } from '@mui/material';
import { useForm } from 'react-hook-form';
import axios from 'axios';

export type MessageCardProps = {
  id: string;
  userID: string;
  text: string;
  createdAt: string;
};

const inputForm: React.FC = () => {
  const { register, handleSubmit } = useForm();

  const handleInput = (data: any) => {
    const formData = new FormData();
    formData.append('user_id', data.user_id);
    formData.append('text', data.text);
    axios
      .post(`http://localhost:8080/post`, formData, {})
      .then(() => {
        console.log(data);
      })
      .catch((e) => {
        alert(`ERROR: HTTP ERROR due to: ${e}`);
      });
  };

  return (
    <Box>
      <form onSubmit={handleSubmit(handleInput)}>
        <input id="user_id" {...register('user_id')} />
        <input id="text" {...register('text')} />
        <button type="submit">送信</button>
      </form>
    </Box>
  );
};

export default inputForm;
