import React from 'react';

interface Props {
    children: React.ReactNode;
}

interface State {
    hasError: boolean;
}

class ErrorBoundary extends React.Component<Props, State> {
    public state: State = {
        hasError: false
    }
    public static getDerivedStateFromError(_: Error): State {
        return { hasError: true }
    }
    public componentDidCatch(error: Error, errorInfo: React.ErrorInfo) {
        console.log("Uncaught error:", error, errorInfo)
    }
    public render() {
        if (this.state.hasError) {
            return <div>RUNTIME ERROR! Please refresh the page.</div>;
        } else {
            return this.props.children;
        }
    }
}

export default ErrorBoundary;