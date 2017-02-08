function contentClass(isShow) {
  if (isShow) {
    return "content";
  }
  return "content invisible";
}

class Toggle extends React.Component {
  constructor(props) {
    super(props);
    this.state = {isShow: false};
    this.handleClick = this.handleClick.bind(this);
  }

  handleClick() {
    this.setState(function(prevState) {
      return {isShow: !prevState.isShow};
    });
  }

  render() {
    return (
      <div>
        <div className='control' onClick={this.handleClick}>Click me to toggle</div>
        <div className={contentClass(this.state.isShow)}>contents</div>
      </div>
    );
  }
}

ReactDOM.render(
  <Toggle />,
  document.getElementById('root')
);
