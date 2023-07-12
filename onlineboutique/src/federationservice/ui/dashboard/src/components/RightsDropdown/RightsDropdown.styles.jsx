import styled from "styled-components";

export const RightsTagTrue = styled.span`
  background-color: #dff0d8; 
  border: 1px solid #d6e9c6;
  border-radius: 4px;
  font-size: 12px;
  margin: 2px;
  padding: 2px 4px;
  display: inline-block;
  color: #3c763d;
  cursor: pointer;
  &:hover {
    opacity: 0.7;
  } 
`;

export const RightsTagFalse = styled.span`
  background-color: #f2dede;
  border: 1px solid #ebccd1;
  border-radius: 4px;
  font-size: 12px;
  margin: 2px;
  padding: 2px 4px;
  display: inline-block;
  color: #a94442;
  cursor: pointer;
  &:hover {
    opacity: 0.7;
  }
`;