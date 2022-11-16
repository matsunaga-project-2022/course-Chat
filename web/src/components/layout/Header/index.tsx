import React from 'react';
import {
  AppBar,
  Avatar,
  Container,
  IconButton,
  Toolbar,
  Typography,
} from '@mui/material';
import MenuIcon from '@mui/icons-material/Menu';

export type HeaderProps = {
  onMenuButtonClicked: React.Dispatch<React.SetStateAction<boolean>>;
};
const drawerWidth = 240;

const Header = (props: HeaderProps) => {
  const { onMenuButtonClicked } = props;
  return (
    <AppBar
      color="transparent"
      position="fixed"
      elevation={0}
      sx={{
        borderBottomStyle: 'solid',
        borderBottomColor: '#EDF2F7',
        borderBottomWidth: '1px',
        width: { sm: `calc(100% - ${drawerWidth}px)` },
        ml: { sm: `${drawerWidth}px` },
      }}
    >
      <Container maxWidth="xl">
        <Toolbar
          sx={{
            display: 'flex',
            mr: 1,
          }}
        >
          <Typography
            variant="h6"
            noWrap
            component="a"
            href="/"
            sx={{
              ml: -4,
              mr: 2,
              fontSize: { xs: '15px', md: '20px' },
              fontWeight: 700,
              letterSpacing: '.3rem',
              color: 'inherit',
              textDecoration: 'none',
            }}
          >
            ウタ
          </Typography>
        </Toolbar>
      </Container>
    </AppBar>
  );
};

export default Header;
