import {useState} from 'react';
import {useDispatch, useSelector} from 'react-redux';
import {addPlayerAction} from './actions/addPlayer';
import {startGameAction} from './actions/startGame';
import {selectRoundNumber} from './selectors/game';
import {selectPlayerCount, selectPlayerNames} from './selectors/players';

function App() {
  const dispatch = useDispatch();
  const [name, updateName] = useState('');
  const roundNumber = useSelector(selectRoundNumber);
  const playerNames = useSelector(selectPlayerNames);
  const playerCount = useSelector(selectPlayerCount);

  return (
    <div>
      {roundNumber > 0 ? (
        <h2>Round {roundNumber}</h2>
      ) : (
        <>
          <h2>Getting Started</h2>
          {playerNames.map((name) => (
            <p key={name}>{name}</p>
          ))}
          <input
            value={name}
            onChange={(event) => updateName(event.currentTarget.value)}></input>
          <button
            onClick={() => {
              dispatch(addPlayerAction(name));
              updateName('');
            }}>
            ADD PLAYER
          </button>
          <div>
            <button onClick={() => dispatch(startGameAction())}>
              START ({playerCount} PLAYERS)
            </button>
          </div>
        </>
      )}
    </div>
  );
}

export default App;
