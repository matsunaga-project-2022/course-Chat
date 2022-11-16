import React from 'react';
import {
  Avatar,
  Divider,
  Drawer,
  List,
  ListItem,
  ListItemAvatar,
  ListItemButton,
  ListItemIcon,
  ListItemText,
  Toolbar,
  Typography,
} from '@mui/material';
import LockIcon from '@mui/icons-material/Lock';
import { useDisplaySize } from '@/utils/useDisplaySize';

const drawerWidth = 240;

export type SideMenuProps = {
  isOpen: boolean;
};

const SideMenu = (props: SideMenuProps) => {
  const { isOpen } = props;
  const { isMobileSize } = useDisplaySize();
  return (
    <Drawer
      sx={{
        width: drawerWidth,
        flexShrink: 0,
        '& .MuiDrawer-paper': {
          width: drawerWidth,
          boxSizing: 'border-box',
        },
      }}
      variant={isMobileSize ? 'persistent' : 'permanent'}
      anchor="left"
      open={isOpen}
    >
      <Toolbar>
        <Typography
          variant="h6"
          noWrap
          component="a"
          href="/"
          sx={{
            mr: 2,
            fontSize: { xs: '0px', md: '15px' },
            fontWeight: 700,
            letterSpacing: '.3rem',
            color: 'inherit',
            textDecoration: 'none',
          }}
        >
          Messages
        </Typography>
      </Toolbar>
      <List>
        {[
          {
            id: '1',
            name: 'ウタ',
            iconURL: 'https://pbs.twimg.com/media/FRta1aTaQAAlq5w.jpg',
          },
          {
            id: '2',
            name: 'ウタ',
            iconURL: 'https://pbs.twimg.com/media/FRta1aTaQAAlq5w.jpg',
          },
          {
            id: '3',
            name: 'ウタ',
            iconURL: 'https://pbs.twimg.com/media/FRta1aTaQAAlq5w.jpg',
          },
          {
            id: '4',
            name: 'ウタ',
            iconURL: 'https://pbs.twimg.com/media/FRta1aTaQAAlq5w.jpg',
          },
        ].map((user, index) => (
          <ListItem
            key={user.id}
            disablePadding
            sx={{
              px: 2,
            }}
          >
            <ListItemButton component="a" href="#simple-list">
              <ListItemAvatar>
                <Avatar alt={user.name} src={user.iconURL} />
              </ListItemAvatar>
              <ListItemText
                primary={user.name}
                secondary={
                  <Typography
                    sx={{ display: 'inline' }}
                    component="span"
                    variant="body2"
                    color="text.secondary"
                  >
                    Ali Connors
                  </Typography>
                }
              />
            </ListItemButton>
          </ListItem>
        ))}
      </List>
    </Drawer>
  );
};

export default SideMenu;
