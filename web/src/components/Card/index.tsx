import React from 'react';
import {
  Card,
  CardHeader,
  CardContent,
  Avatar,
  Typography,
} from '@mui/material';
import { red } from '@mui/material/colors';

export type MessageCardProps = {
  id: string;
  userID: string;
  text: string;
  createdAt: string;
};

export const MessageCard = (props: MessageCardProps) => {
  const { id, userID, text, createdAt } = props;

  return (
    <Card key={id} sx={{ width: '100%' }} elevation={0}>
      <CardHeader
        avatar={
          <Avatar sx={{ bgcolor: red[500] }} aria-label="recipe">
            F
          </Avatar>
        }
        title={userID}
        subheader={createdAt}
      />
      <CardContent>
        <Typography variant="body2" color="text.secondary">
          {text}
        </Typography>
      </CardContent>
    </Card>
  );
};
