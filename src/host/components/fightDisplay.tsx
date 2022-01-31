import {useDispatch, useSelector} from 'react-redux';
import {FightStatus} from '../../models';
import {getCombatantsAction} from '../actions/getCombatants';
import {progressFightAction} from '../actions/progressFight';
import {resetFightAction} from '../actions/resetFight';
import {
  selectFightStatus,
  selectFightStatusDescription,
  selectFightStepDetails,
} from '../selectors/fight';

function FightDisplay() {
  const combatants = useSelector(selectFightStepDetails);
  const status = useSelector(selectFightStatus);
  const statusString = useSelector(selectFightStatusDescription);

  const dispatch = useDispatch();

  if (combatants.length === 0) {
    dispatch(getCombatantsAction());
  }

  return (
    <div>
      <h3>Fight</h3>
      <p>{statusString}</p>
      {combatants.map((c) => (
        <p>
          <b>{c.name}</b> {c.currentHealth}/{c.maxHealth}
        </p>
      ))}
      {status === FightStatus.Finished ? (
        <button onClick={() => dispatch(resetFightAction())}>RESET</button>
      ) : (
        <button onClick={() => dispatch(progressFightAction())}>NEXT</button>
      )}
    </div>
  );
}

export default FightDisplay;
