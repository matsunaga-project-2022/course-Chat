import React from 'react';
import {
  Box,
  Divider,
  Drawer,
  List,
  ListItem,
  ListItemButton,
  ListItemIcon,
  ListItemText,
  Toolbar,
} from '@mui/material';
import InboxIcon from '@mui/icons-material/MoveToInbox';
import MailIcon from '@mui/icons-material/Mail';
import SideMenu from '@/components/layout/SideMenu';
import Header from '@/components/layout/Header';
import { Outlet } from 'react-router-dom';

const drawerWidth = 240;

const Layout: React.FC = () => {
  const [open, setOpen] = React.useState<boolean>(false);

  return (
    <Box sx={{ display: 'flex' }}>
      <Header onMenuButtonClicked={setOpen} />
      <SideMenu isOpen={open} />
      <Outlet />
    </Box>
  );
};

export default Layout;
