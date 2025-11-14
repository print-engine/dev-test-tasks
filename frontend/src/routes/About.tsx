import { Link } from 'react-router-dom';

function About() {
  return (
    <div>
      <h1>About Page</h1>
      <p>This is a Vite + React + TypeScript template with:</p>
      <ul>
        <li>React Router for navigation</li>
        <li>Redux Toolkit for state management</li>
        <li>Redux Saga for side effects</li>
        <li>ESLint + Prettier for code quality</li>
        <li>Vitest for testing</li>
      </ul>

      <nav>
        <Link to="/">Go to Home</Link>
      </nav>
    </div>
  );
}

export default About;
