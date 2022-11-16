import React from 'react';
import {
  Avatar,
  Box,
  Card,
  CardActions,
  CardContent,
  CardHeader,
  CardMedia,
  IconButton,
  List,
  ListItem,
  ListItemText,
  Typography,
} from '@mui/material';
import { grey, red } from '@mui/material/colors';

const Home: React.FC = () => (
  <Box
    sx={{
      marginTop: '60px',
    }}
  >
    <Card sx={{ width: '100%' }} elevation={0}>
      <CardHeader
        avatar={
          <Avatar sx={{ bgcolor: red[500] }} aria-label="recipe">
            F
          </Avatar>
        }
        title="世界の歌姫 ウタ"
        subheader="16:00"
      />
      <CardContent>
        <Typography variant="body2" color="text.secondary">
          でもだからこそ、正解がないことにより広がる無限の荒野、そこを冒険して自分なりのアプローチをしていく世界観が非常に面白い。そんな無限大の可能性を秘めた領域の中で、貴方はプロダクトの価値を高めるという目的のもと旅を続ける。「そんなことする意味なんてあるの？」なんて言う外野からの罵倒という名の隕石を浴びながら、果てなき道を歩き続ける。歩く先は闇か、炎上か、それとも光り輝く鉱石(功績)か。ようこそ、アーキテクトの世界へ。
        </Typography>
      </CardContent>
    </Card>
  </Box>
);

export default Home;
