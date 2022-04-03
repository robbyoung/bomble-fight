import {useSelector} from 'react-redux';
import NewBetForm from './components/newBetForm';
import NewPlayerForm from './components/newPlayerForm';
import {selectBetExists} from './selectors/bet';
import {selectPlayer} from './selectors/player';
import FightWaiter from './components/fightWaiter';
import {selectFightStatus} from './selectors/fight';
import {FightStatus} from '../models';

function App() {
  const player = useSelector(selectPlayer);
  const hasBet = useSelector(selectBetExists);
  const fightStatus = useSelector(selectFightStatus);

  if (player.name === '') {
    return <NewPlayerForm />;
  }

  if (!hasBet && fightStatus === FightStatus.Pending) {
    return <NewBetForm />;
  }

  return <FightWaiter />;
}

export default App;
