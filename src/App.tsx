import {useDispatch, useSelector} from 'react-redux';
import {startGameAction} from './actions/startGame';
import {getRoundNumber} from './selectors/game';

function App() {
  const dispatch = useDispatch();
  const roundNumber = useSelector(getRoundNumber);

  return (
    <div>
      {roundNumber > 0 ? <h2>Round {roundNumber}</h2> : undefined}
      <h1>Press button to start a new tournament</h1>
      <button onClick={() => dispatch(startGameAction(2))}>START</button>
    </div>
  );
}

export default App;
