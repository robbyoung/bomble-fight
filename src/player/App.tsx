import {useSelector} from 'react-redux';
import {selectFightStatusDescription} from './selectors/fight';
import NewBetForm from './components/newBetForm';
import NewPlayerForm from './components/newPlayerForm';
import {selectBetExists} from './selectors/bet';
import {selectPlayer} from './selectors/player';

function App() {
  const player = useSelector(selectPlayer);
  const hasBet = useSelector(selectBetExists);
  const fightStatus = useSelector(selectFightStatusDescription);

  return (
    <div>
      {player.name === '' ? <NewPlayerForm /> : null}
      {player.name !== '' && !hasBet ? <NewBetForm /> : null}
      {hasBet ? <p>{fightStatus}</p> : null}
    </div>
  );
}

export default App;
