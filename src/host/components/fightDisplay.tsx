import {useDispatch, useSelector} from 'react-redux';
import {getCombatantsAction} from '../actions/getCombatants';
import {progressFightAction} from '../actions/progressFight';
import {
  selectFightStatusDescription,
  selectFightStepDetails,
} from '../selectors/fight';

function FightDisplay() {
  const combatants = useSelector(selectFightStepDetails);
  const status = useSelector(selectFightStatusDescription);

  const dispatch = useDispatch();

  if (combatants.length === 0) {
    dispatch(getCombatantsAction());
  }

  return (
    <div>
      <h3>Fight</h3>
      <p>{status}</p>
      {combatants.map((c) => (
        <p>
          <b>{c.name}</b> {c.currentHealth}/{c.maxHealth}
        </p>
      ))}
      <button onClick={() => dispatch(progressFightAction())}>NEXT</button>
    </div>
  );
}

export default FightDisplay;
