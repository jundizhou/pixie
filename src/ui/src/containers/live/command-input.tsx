import Autocomplete from 'components/autocomplete/autocomplete';
import { CompletionHeader, CompletionItem } from 'components/autocomplete/completions';
import MagicIcon from 'components/icons/magic';
import * as React from 'react';
import { GetPxScripts, Script } from 'utils/script-bundle';

import { createStyles, makeStyles, Theme } from '@material-ui/core';
import Card from '@material-ui/core/Card';
import Modal from '@material-ui/core/Modal';

import { LiveContext } from './context';
import { parseVis } from './vis';

interface CommandInputProps {
  open: boolean;
  onClose: () => void;
}

const useStyles = makeStyles((theme: Theme) =>
  createStyles({
    card: {
      position: 'absolute',
      width: '760px',
      top: '40%',
      left: '50%',
      transform: 'translate(-50%, -20vh)',
    },
    input: {
      maxHeight: '60vh',
    },
  }),
);

const MOCK_COMMANDS = [
  'px/service_info',
  'px/service_stats',
  'px/network_stats',
  'px/namespace_stats',
  'service_info',
  'network_info',
  'cpu_info',
  'memory_info',
  'service_info',
  'network_info',
  'cpu_info',
  'memory_info',
];

const DESCRIPTION = `
  Lorem ipsum dolor sit amet, consectetur adipiscing elit. Vestibulum aliquet tincidunt ligula. \
  Vivamus congue odio sed nisi volutpat venenatis. Morbi pulvinar neque nisl, in ornare massa \
  pellentesque viverra. Pellentesque auctor suscipit ex, quis dignissim ligula. Mauris urna purus, \
`;

// TODO(malthus): Figure out the lifecycle of this component. When a command is selected,
// should the component clear the input? What about when the input is dismised?

const CommandInput: React.FC<CommandInputProps> = ({ open, onClose }) => {
  const classes = useStyles();

  const [scriptsMap, setScriptsMap] = React.useState<Map<string, Script>>(null);
  const [completions, setCompletions] = React.useState<CompletionItem[]>([]);

  React.useEffect(() => {
    GetPxScripts().then((examples) => {
      setCompletions(examples.map((s) => ({
        type: 'item',
        id: s.id,
        title: s.id,
        description: s.description,
      })));
      setScriptsMap(new Map(examples.map((s) => [s.id, s])));
    });
  }, []);

  const { setScripts, executeScript } = React.useContext(LiveContext);

  const getCompletions = React.useCallback((input) => {
    if (!input) {
      return Promise.resolve([{ type: 'header', header: 'Example Scripts' } as CompletionHeader, ...completions]);
    }
    return Promise.resolve(completions.filter((completion) => completion.title.includes(input)));
  }, [completions]);

  const selectScript = React.useCallback((id) => {
    const script = scriptsMap.get(id);
    if (script) {
      setScripts(script.code, script.vis, { title: script.title, id: script.id }, {});
      executeScript(script.code, parseVis(script.vis));
    }
    onClose();
  }, [scriptsMap, executeScript]);

  return (
    <Modal open={open} onClose={onClose} BackdropProps={{}}>
      <Card className={classes.card}>
        <Autocomplete
          className={classes.input}
          placeholder='Pixie Command'
          prefix={<MagicIcon />}
          onSelection={selectScript}
          getCompletions={getCompletions}
        />
      </Card>
    </Modal>
  );
};

export default CommandInput;
