import { Link } from 'react-router-dom';
import { useAppDispatch, useAppSelector } from '../store/hooks';
import { increment, decrement, reset, selectCount } from '../features/counter/counterSlice';

function Home() {
  const count = useAppSelector(selectCount);
  const dispatch = useAppDispatch();

  return (
    <div>
      <h1>Home Page</h1>
      <p>Welcome to the frontend template!</p>

      <div style={{ margin: '20px 0' }}>
        <h2>Redux Counter Example</h2>
        <div style={{ fontSize: '2rem', margin: '10px 0' }}>{count}</div>
        <div style={{ display: 'flex', gap: '10px' }}>
          <button onClick={() => dispatch(decrement())}>-</button>
          <button onClick={() => dispatch(increment())}>+</button>
          <button onClick={() => dispatch(reset())}>Reset</button>
        </div>
      </div>

      <nav>
        <Link to="/about">Go to About</Link>
      </nav>
    </div>
  );
}

export default Home;
