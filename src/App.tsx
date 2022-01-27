import {useState} from 'react';
import {useDispatch, useSelector} from 'react-redux';
import {addBetAction} from './actions/addBet';
import {addPlayerAction} from './actions/addPlayer';
import {getCombatantsAction} from './actions/getCombatants';
import {getPlayersAction} from './actions/getPlayers';
import {progressFightAction} from './actions/progressFight';
import {startGameAction} from './actions/startGame';
import {selectBetsByPlayer} from './selectors/bets';
import {selectCombatantNames, selectCombatants} from './selectors/combatants';
import {selectFightStatus} from './selectors/fight';
import {selectRoundNumber} from './selectors/game';
import {
  selectPlayerCount,
  selectPlayerNames,
  selectPlayers,
} from './selectors/players';

function App() {
  const dispatch = useDispatch();
  const [name, updateName] = useState('');
  const roundNumber = useSelector(selectRoundNumber);
  const playerNames = useSelector(selectPlayerNames);
  const playerCount = useSelector(selectPlayerCount);
  const bets = useSelector(selectBetsByPlayer);
  const players = useSelector(selectPlayers);
  const combatants = useSelector(selectCombatants);
  const combatantNames = useSelector(selectCombatantNames);
  const fightStatus = useSelector(selectFightStatus);

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
          <button onClick={() => dispatch(getPlayersAction())}>
            REFRESH PLAYERS
          </button>
        </>
      )}
      <div>
        <h3>Combatants</h3>
        {combatantNames.map((name) => (
          <p key={name}>{name}</p>
        ))}
        <button onClick={() => dispatch(getCombatantsAction())}>
          REFRESH COMBATANTS
        </button>
      </div>
      <div>
        <h3>Bets</h3>
        {bets.map((b) => (
          <p key={b.playerName}>
            {b.playerName}: {b.combatantName}
          </p>
        ))}
        <button
          onClick={() =>
            dispatch(addBetAction(players[0].id, combatants[0].id, 100))
          }>
          BET
        </button>
      </div>
      <div>
        <h3>Fight</h3>
        <p>Fight status: {fightStatus}</p>
        <button onClick={() => dispatch(progressFightAction())}>
          PROGRESS
        </button>
      </div>
    </div>
  );
}

export default App;
