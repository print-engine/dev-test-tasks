import { all } from 'redux-saga/effects';

// Example saga - you can add more sagas here
// function* exampleSaga() {
//   yield takeEvery('ACTION_TYPE', workerSaga);
// }

export default function* rootSaga() {
  yield all([
    // Add your sagas here
    // exampleSaga(),
  ]);
}
