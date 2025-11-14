import { describe, it, expect } from 'vitest';
import { render, screen } from '@testing-library/react';
import { Provider } from 'react-redux';
import { BrowserRouter } from 'react-router-dom';
import { configureStore } from '@reduxjs/toolkit';
import Home from './Home';
import counterReducer from '../features/counter/counterSlice';

const createTestStore = () => {
  return configureStore({
    reducer: {
      counter: counterReducer,
    },
  });
};

const renderWithProviders = (component: React.ReactElement) => {
  const store = createTestStore();
  return render(
    <Provider store={store}>
      <BrowserRouter>{component}</BrowserRouter>
    </Provider>
  );
};

describe('Home', () => {
  it('renders home page heading', () => {
    renderWithProviders(<Home />);
    expect(screen.getByText('Home Page')).toBeInTheDocument();
  });

  it('renders counter with initial value of 0', () => {
    renderWithProviders(<Home />);
    expect(screen.getByText('0')).toBeInTheDocument();
  });

  it('renders navigation link to about page', () => {
    renderWithProviders(<Home />);
    const link = screen.getByText('Go to About');
    expect(link).toBeInTheDocument();
    expect(link).toHaveAttribute('href', '/about');
  });
});
