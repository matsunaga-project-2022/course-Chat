import { Theme, useMediaQuery } from '@mui/material';

export const useDisplaySize = () => {
  const isMobileSize = useMediaQuery((theme: Theme) =>
    theme.breakpoints.down('xs')
  );
  return { isMobileSize };
};
