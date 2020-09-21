import React from 'react';
import { Link } from 'react-router-dom';
import { FormContainer } from '../components/FormContainer';

export const NotFound = () => {
  return (
    <FormContainer>
      <h1>ページが見つかりません。</h1>
      <Link to="/">HOME</Link>
    </FormContainer>
  );
};
