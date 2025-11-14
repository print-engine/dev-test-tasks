import { describe, it, expect } from 'vitest';
import counterReducer, {
  increment,
  decrement,
  incrementByAmount,
  reset,
} from './counterSlice';

describe('counter reducer', () => {
  const initialState = {
    value: 0,
    status: 'idle' as const,
  };

  it('should handle initial state', () => {
    expect(counterReducer(undefined, { type: 'unknown' })).toEqual(initialState);
  });

  it('should handle increment', () => {
    const actual = counterReducer(initialState, increment());
    expect(actual.value).toEqual(1);
  });

  it('should handle decrement', () => {
    const actual = counterReducer(initialState, decrement());
    expect(actual.value).toEqual(-1);
  });

  it('should handle incrementByAmount', () => {
    const actual = counterReducer(initialState, incrementByAmount(5));
    expect(actual.value).toEqual(5);
  });

  it('should handle reset', () => {
    const state = { value: 10, status: 'idle' as const };
    const actual = counterReducer(state, reset());
    expect(actual.value).toEqual(0);
  });
});
