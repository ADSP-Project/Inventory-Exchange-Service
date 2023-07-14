import styled from 'styled-components';

export const Table = styled.table`
  width: 100%;
  border-collapse: collapse;

  th,
  td {
    border: 1px solid #ddd;
    padding: 0.5em;
  }

  th {
    background-color: #f5f5f5;
  }
`;

export const Button = styled.button`
  background-color: #007bff;
  color: white;
  border: none;
  padding: 0.5em 1em;
  border-radius: 0.25em;
  cursor: pointer;
  margin-right: 0.5em;
`;

export const RightsTagTrue = styled.span`
  background-color: #dff0d8; 
  border: 1px solid #d6e9c6;
  border-radius: 4px;
  font-size: 12px;
  margin: 2px;
  padding: 2px 4px;
  display: inline-block;
  color: #3c763d; 
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
`;