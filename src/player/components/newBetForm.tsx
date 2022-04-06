import {useEffect, useState} from 'react';
import {useDispatch, useSelector} from 'react-redux';
import {getCombatantsAction} from '../../host/actions/getCombatants';
import {addBetAction} from '../actions/addBet';
import {selectCombatants} from '../selectors/combatants';
import {selectPlayer} from '../selectors/player';
import './newBetForm.css';

function NewBetForm() {
  const [combatantId, setCombatantId] = useState('');
  const [amount, setAmount] = useState(0);
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
          <button
            className={
              c.id === combatantId ? 'combatant selected' : 'combatant'
            }
            onClick={() => setCombatantId(c.id)}>
            {c.name}
          </button>
        ))}
      </div>
      <p>Your money: ${player.money}</p>
      <div>
        <span>Amount: </span>
        <input
          type="number"
          onChange={(event) => setAmount(parseInt(event.target.value))}></input>
      </div>
      <button
        onClick={() => dispatch(addBetAction(player.id, combatantId, amount))}>
        SUBMIT
      </button>
    </div>
  );
}

export default NewBetForm;
