import {useDispatch, useSelector} from 'react-redux';
import {
  startPlayersPollAction,
  stopPlayersPollAction,
} from '../actions/getPlayers';
import {selectPlayers} from '../selectors/players';

function PlayerList() {
  const players = useSelector(selectPlayers);
  const dispatch = useDispatch();

  if (players.length === 0) {
    dispatch(startPlayersPollAction());
  }

  return (
    <div>
      <h3>Players</h3>
      {players.map((p) => (
        <p>
          <b>{p.name}</b> ${p.money}
        </p>
      ))}
      <button onClick={() => dispatch(stopPlayersPollAction())}>READY</button>
    </div>
  );
}

export default PlayerList;
