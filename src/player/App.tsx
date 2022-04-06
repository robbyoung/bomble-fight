import {useDispatch, useSelector} from 'react-redux';
import NewBetForm from './components/newBetForm';
import NewPlayerForm from './components/newPlayerForm';
import {selectBetExists} from './selectors/bet';
import {selectPlayer} from './selectors/player';
import FightWaiter from './components/fightWaiter';
import {selectFightStatus} from './selectors/fight';
import {FightStatus} from '../models';
import {useState} from 'react';
import {
  startPlayerStatePollAction,
  stopPlayerStatePollAction,
} from './actions/getPlayerState';

function App() {
  const player = useSelector(selectPlayer);
  const hasBet = useSelector(selectBetExists);
  const fightStatus = useSelector(selectFightStatus);
  const dispatch = useDispatch();

  const [isPollingState, setIsPollingState] = useState(false);

  if (player.name === '') {
    return <NewPlayerForm />;
  }

  if (!hasBet && fightStatus === FightStatus.Pending) {
    if (isPollingState) {
      dispatch(stopPlayerStatePollAction());
      setIsPollingState(false);
    }
    return <NewBetForm />;
  }

  if (!isPollingState) {
    dispatch(startPlayerStatePollAction());
    setIsPollingState(true);
  }
  return <FightWaiter />;
}

export default App;
