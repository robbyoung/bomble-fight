import {useState} from 'react';
import {useDispatch} from 'react-redux';
import {addPlayerAction} from '../actions/addPlayer';

function NewPlayerForm() {
  const [name, setName] = useState('');
  const dispatch = useDispatch();

  return (
    <div>
      <p>Player name:</p>
      <input
        onChange={(event) => setName(event.target.value)}
        value={name}></input>
      <button onClick={() => dispatch(addPlayerAction(name))}>SUBMIT</button>
    </div>
  );
}

export default NewPlayerForm;
