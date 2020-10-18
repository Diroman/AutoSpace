import * as React from 'react';

import { createGlobalStyle } from 'styled-components';

export const ThemeProvider = ({ children }) => {

    return (
        <>
            <GlobalStyles />
            {children}
        </>
    );
};

const GlobalStyles = createGlobalStyle`
  html {
    height: 100%;
  }
  body {
    background: #eeeeee;
    font-size: 1.6rem;
    height: 100%;
    overflow: scroll;
    width: 100%;
  }
  #app {
    height: 100%;
  }
  *::-webkit-scrollbar {
    display: none;
  }
  * {
    -ms-overflow-style: none;
    overflow: -moz-scrollbars-none;
    scrollbar-width: none;
  }
`;
