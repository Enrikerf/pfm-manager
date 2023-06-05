package TaskTest

/*func TestLoad(t *testing.T) {
	type args struct {
		id                Task.Id
		host              Host.Vo
		port              Port.Vo
		stepVos           []Step.Vo
		communicationMode CommunicationMode.Mode
		executionMode     ExecutionMode.Mode
		status            Status.Status
	}
	tests := []struct {
		name    string
		args    args
		want    Task.Task
		wantErr bool
	}{
		{
			name: "",
			args: args{
				id:                nil,
				host:              nil,
				port:              nil,
				stepVos:           nil,
				communicationMode: nil,
				executionMode:     nil,
				status:            nil,
			},
			want:    nil,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Task.Load(tt.args.id, tt.args.host, tt.args.port, tt.args.stepVos, tt.args.communicationMode, tt.args.executionMode, tt.args.status)
			if (err != nil) != tt.wantErr {
				t.Errorf("Load() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Load() got = %v, want %v", got, tt.want)
			}
		})
	}
}*/

/*func TestNew(t *testing.T) {
	type args struct {
		host              Host.Vo
		port              Port.Vo
		stepVos           []Step.Vo
		communicationMode CommunicationMode.Mode
		executionMode     ExecutionMode.Mode
	}
	tests := []struct {
		name    string
		args    args
		want    Task.Task
		wantErr bool
	}{
		{
			name: "",
			args: args{
				host:              HostTest.NewVoMock("0.0.0.0"),
				port:              nil,
				stepVos:           nil,
				communicationMode: nil,
				executionMode:     nil,
			},
			want:    nil,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Task.New(tt.args.host, tt.args.port, tt.args.stepVos, tt.args.communicationMode, tt.args.executionMode)
			if (err != nil) != tt.wantErr {
				t.Errorf("New() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("New() got = %v, want %v", got, tt.want)
			}
		})
	}
}*/
