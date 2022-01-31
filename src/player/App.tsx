import {useSelector} from 'react-redux';
import NewBetForm from './components/newBetForm';
import NewPlayerForm from './components/newPlayerForm';
import {selectPlayer} from './selectors/player';

function App() {
  const player = useSelector(selectPlayer);

  return <div>{player.name === '' ? <NewPlayerForm /> : <NewBetForm />}</div>;
}

export default App;
