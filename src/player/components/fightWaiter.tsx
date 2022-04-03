import {useDispatch, useSelector} from 'react-redux';
import {FightStatus} from '../../models';
import {getPlayerState} from '../actions/getPlayerState';
import {selectFightStatus} from '../selectors/fight';
import {selectPlayer} from '../selectors/player';

function FightWaiter() {
  const status = useSelector(selectFightStatus);
  const player = useSelector(selectPlayer);
  const dispatch = useDispatch();

  var message;
  switch (status) {
    case FightStatus.Pending:
    case FightStatus.Starting:
      message = "Fight hasn't started yet.";
      break;
    case FightStatus.Active:
      message = 'Fight is happening now.';
      break;
    case FightStatus.Finished:
      message = `Fight has concluded! You have $${player.money} left.`;
      break;
  }

  return (
    <div>
      <p>{message}</p>
      <button onClick={() => dispatch(getPlayerState(player.id))}>
        REFRESH
      </button>
    </div>
  );
}

export default FightWaiter;
