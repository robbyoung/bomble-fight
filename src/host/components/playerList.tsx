import {useDispatch, useSelector} from 'react-redux';
import {getPlayersAction} from '../actions/getPlayers';
import {selectPlayers} from '../selectors/players';

function PlayerList() {
  const players = useSelector(selectPlayers);
  const dispatch = useDispatch();

  return (
    <div>
      <h3>Players</h3>
      {players.map((p) => (
        <p>
          <b>{p.name}</b> ${p.money}
        </p>
      ))}
      <button onClick={() => dispatch(getPlayersAction())}>REFRESH</button>
    </div>
  );
}

export default PlayerList;
