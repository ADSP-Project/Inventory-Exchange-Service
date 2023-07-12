import styled from 'styled-components';

export const Tile = styled.div`
  border: 1px solid #ddd;
  border-radius: 4px;
  margin: 10px;
  padding: 20px;
  max-width: 300px;
`;

export const TileHeader = styled.h2`
  margin: 0;
  margin-bottom: 10px;
  color: #333;
`;

export const TileBody = styled.p`
  margin: 0;
  color: #666;
`;

export const TileFooter = styled.div`
  margin-top: 20px;
`;

export const JoinButton = styled.button`
  background-color: #008CBA; 
  border: none;
  color: white;
  padding: 15px 32px;
  text-align: center;
  text-decoration: none;
  display: inline-block;
  font-size: 16px;
  margin: 4px 2px;
  cursor: pointer;

  &:disabled {
    background-color: gray; 
    cursor: not-allowed;
  }
`;

export const Image = styled.img`
  width: 100%;
  height: 300px;  
  object-fit: cover;
  object-position: center;
`;
