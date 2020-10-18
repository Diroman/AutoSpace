import styled from 'styled-components';
import { Flex } from '../common/components/Flex';
import * as React from 'react';

export const Layout = ({ children }) =>
    <LayoutStyled>
        <ContentStyled>
            {children}
        </ContentStyled>
    </LayoutStyled>
;

const LayoutStyled = styled(Flex)`
  height: 100%;
  overflow: auto;
`;

const ContentStyled = styled.div`
  flex-grow: 1;
`;
