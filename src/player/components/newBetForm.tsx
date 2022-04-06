import {useEffect, useState} from 'react';
import {useDispatch, useSelector} from 'react-redux';
import {getCombatantsAction} from '../../host/actions/getCombatants';
import {addBetAction} from '../actions/addBet';
import {selectCombatants} from '../selectors/combatants';
import {selectPlayer} from '../selectors/player';

function NewBetForm() {
  const [combatantId, setCombatantId] = useState('');
  const combatants = useSelector(selectCombatants);
  const player = useSelector(selectPlayer);
  const dispatch = useDispatch();

  useEffect(() => {
    dispatch(getCombatantsAction());
  }, [dispatch]);

  return (
    <div>
      <div>
        {combatants.map((c) => (
          <button onClick={() => setCombatantId(c.id)}>{c.name}</button>
        ))}
      </div>
      <button
        onClick={() => dispatch(addBetAction(player.id, combatantId, 50))}>
        SUBMIT
      </button>
    </div>
  );
}

export default NewBetForm;
