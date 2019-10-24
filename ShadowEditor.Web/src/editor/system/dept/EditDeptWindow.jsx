import { PropTypes, Window, Content, Buttons, Form, FormControl, Label, Input, Button } from '../../../third_party';

/**
 * 组织机构编辑窗口
 * @author tengge / https://github.com/tengge1
 */
class EditDeptWindow extends React.Component {
    constructor(props) {
        super(props);

        this.state = {
            id: props.id,
            name: props.name,
        };

        this.handleChange = this.handleChange.bind(this);
        this.handleSave = this.handleSave.bind(this, props.callback);
        this.handleClose = this.handleClose.bind(this);
    }

    render() {
        const { id, name } = this.state;

        return <Window
            className={_t('EditDeptWindow')}
            title={id ? _t('Add Department') : _t('Edit Department')}
            style={{ width: '320px', height: '200px' }}
            mask={false}
            onClose={this.handleClose}>
            <Content>
                <Form>
                    <FormControl>
                        <Label>{_t('Name')}</Label>
                        <Input name={'name'} value={name} onChange={this.handleChange}></Input>
                    </FormControl>
                </Form>
            </Content>
            <Buttons>
                <Button onClick={this.handleSave}>{_t('OK')}</Button>
                <Button onClick={this.handleClose}>{_t('Cancel')}</Button>
            </Buttons>
        </Window>;
    }

    handleChange(value, name) {
        this.setState({
            [name]: value,
        });
    }

    handleSave(callback) {
        const { id, name } = this.state;

        if (!name || name.trim() === '') {
            app.toast(_t('Name is not allowed to be empty.'));
            return;
        }

        const url = !id ? `/api/Department/Add` : `/api/Department/Edit`;

        fetch(`${app.options.server}${url}`, {
            method: 'POST',
            headers: {
                'Content-Type': 'application/x-www-form-urlencoded',
            },
            body: `ID=${id}&ParentID=&Name=${name}&AdministratorID=`,
        }).then(response => {
            response.json().then(json => {
                if (json.Code !== 200) {
                    app.toast(_t(json.Msg));
                    return;
                }
                this.handleClose();
                callback && callback();
            });
        });
    }

    handleClose() {
        app.removeElement(this);
    }
}

EditDeptWindow.propTypes = {
    id: PropTypes.string,
    name: PropTypes.string,
    callback: PropTypes.func,
};

EditDeptWindow.defaultProps = {
    id: '',
    name: '',
    callback: null,
};

export default EditDeptWindow;