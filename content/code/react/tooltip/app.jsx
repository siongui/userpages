function Tooltip(props) {
  if (props.isVisible) {
    return <div className="tooltip" style={props.style}>{props.textContent}</div>;
  }
  return <div className="tooltip invisible" style={props.style}>{props.textContent}</div>;
}

class TooltipApp extends React.Component {
  constructor(props) {
    super(props);
    this.state = {
      isShowTooltip: false,
      tooltipContent: "",
      tooltipStyle: {}
    };
    this.ShowTooltip = this.ShowTooltip.bind(this);
    this.HideTooltip = this.HideTooltip.bind(this);
  }

  ShowTooltip(event) {
    var elm = event.target;
    var tTop = (elm.offsetTop + elm.offsetHeight + 5) + 'px';
    var tLeft = elm.offsetLeft + 'px';
    this.setState({
      isShowTooltip: true,
      tooltipContent: event.target.dataset.descr,
      tooltipStyle: {
        top: tTop,
        left: tLeft
      }
    });
  }

  HideTooltip() {
    this.setState({isShowTooltip: false});
  }

  render() {
    return (
      <div>
        <Tooltip style={this.state.tooltipStyle} textContent={this.state.tooltipContent} isVisible={this.state.isShowTooltip} />
        <p>This is a example of{" "}
          <span onMouseEnter={this.ShowTooltip} onMouseLeave={this.HideTooltip} data-descr="Hello, I am tooltip!">tooltip</span> via{" "}
          <span onMouseEnter={this.ShowTooltip} onMouseLeave={this.HideTooltip} data-descr="Hello, I am React!">React</span></p>
      </div>
    );
  }
}

ReactDOM.render(
  <TooltipApp />,
  document.getElementById('root')
);
