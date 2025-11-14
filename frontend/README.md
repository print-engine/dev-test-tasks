# Frontend - React TypeScript Template

A modern React TypeScript template built with Vite, featuring Redux Toolkit, Redux Saga, React Router, and comprehensive tooling.

## Features

- **React 19** with TypeScript for type-safe development
- **Vite** for fast development and optimized builds
- **React Router** for client-side routing
- **Redux Toolkit** for state management
- **Redux Saga** for handling side effects
- **ESLint + Prettier** for code quality and formatting
- **Vitest** with React Testing Library for unit and integration tests
- Hot Module Replacement (HMR) for instant feedback during development

## Prerequisites

- Node.js 18+ and npm

## Getting Started

### Installation

```bash
npm install
```

### Development

Start the development server with hot reload:

```bash
npm run dev
```

The app will be available at `http://localhost:5173`

### Building for Production

```bash
npm run build
```

Build output will be in the `dist/` directory.

### Preview Production Build

```bash
npm run preview
```

## Available Scripts

- `npm run dev` - Start development server
- `npm run build` - Build for production
- `npm test` - Run tests in watch mode
- `npm test -- --run` - Run tests once
- `npm run lint` - Lint code with ESLint
- `npm run format` - Format code with Prettier
- `npm run preview` - Preview production build locally

## Project Structure

```
frontend/
├── src/
│   ├── features/           # Feature-based modules
│   │   └── counter/        # Example counter feature
│   │       ├── counterSlice.ts       # Redux slice
│   │       └── counterSlice.test.ts  # Unit tests
│   ├── routes/             # Route components
│   │   ├── Home.tsx        # Home page
│   │   ├── Home.test.tsx   # Home page tests
│   │   └── About.tsx       # About page
│   ├── store/              # Redux store configuration
│   │   ├── store.ts        # Store setup with saga middleware
│   │   ├── hooks.ts        # Typed Redux hooks
│   │   └── sagas.ts        # Root saga
│   ├── test/               # Test configuration
│   │   └── setup.ts        # Test setup file
│   ├── App.tsx             # Main app component with routing
│   ├── main.tsx            # App entry point
│   └── index.css           # Global styles
├── .prettierrc             # Prettier configuration
├── eslint.config.js        # ESLint configuration
├── vitest.config.ts        # Vitest configuration
├── tsconfig.json           # TypeScript configuration
├── vite.config.ts          # Vite configuration
└── package.json            # Dependencies and scripts
```

## Tech Stack Details

### State Management

- **Redux Toolkit**: Simplified Redux with less boilerplate
- **Redux Saga**: Handle complex async flows and side effects
- Typed hooks (`useAppDispatch`, `useAppSelector`) for type safety

Example usage:

```typescript
import { useAppDispatch, useAppSelector } from '../store/hooks';
import { increment, selectCount } from '../features/counter/counterSlice';

function MyComponent() {
  const count = useAppSelector(selectCount);
  const dispatch = useAppDispatch();

  return <button onClick={() => dispatch(increment())}>{count}</button>;
}
```

### Routing

React Router v7 with declarative routing:

```typescript
<Routes>
  <Route path="/" element={<Home />} />
  <Route path="/about" element={<About />} />
</Routes>
```

### Testing

- **Vitest**: Fast unit test framework
- **React Testing Library**: Component testing utilities
- **@testing-library/jest-dom**: Custom matchers for assertions

Run tests:

```bash
npm test
```

### Code Quality

- **ESLint**: Configured with TypeScript, React, and Prettier rules
- **Prettier**: Consistent code formatting
- Pre-configured rules for React hooks and modern patterns

## Development Guidelines

### Adding New Features

1. Create a new directory in `src/features/`
2. Add slice file with Redux Toolkit's `createSlice`
3. Register reducer in `src/store/store.ts`
4. Add sagas to `src/store/sagas.ts` if needed
5. Write tests for your slice and components

### Adding New Routes

1. Create component in `src/routes/`
2. Add route in `src/App.tsx`
3. Write tests for the route component

### Adding Sagas

Example saga for async operations:

```typescript
import { call, put, takeEvery } from 'redux-saga/effects';

function* fetchDataSaga() {
  try {
    const data = yield call(api.fetchData);
    yield put(dataLoadedSuccess(data));
  } catch (error) {
    yield put(dataLoadedError(error));
  }
}

export function* watchFetchData() {
  yield takeEvery('data/fetchRequested', fetchDataSaga);
}
```

## Best Practices Demonstrated

- Feature-based project structure
- Typed Redux hooks for type safety
- Separation of concerns (routes, features, store)
- Comprehensive test coverage
- Code quality tools (ESLint + Prettier)
- Modern React patterns (hooks, functional components)
- Type-safe imports with TypeScript strict mode

## Configuration Files

- `.prettierrc` - Prettier formatting rules
- `eslint.config.js` - ESLint linting rules
- `vitest.config.ts` - Test runner configuration
- `tsconfig.json` - TypeScript compiler options
- `vite.config.ts` - Build tool configuration

## Browser Support

Modern browsers with ES2020+ support. Configure targets in `vite.config.ts` if needed.

## License

This is a template project for educational purposes.
