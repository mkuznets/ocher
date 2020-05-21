import os
import shutil
import sys
import tempfile

from grpc_tools import protoc

if __name__ == '__main__':

    proto_paths = sys.argv[1:]
    if not proto_paths:
        print("expected .proto filenames", file=sys.stderr)
        sys.exit(1)

    with tempfile.TemporaryDirectory() as wd:
        worker_dir = os.path.join(wd, 'ocher/pb')

        os.makedirs(worker_dir, exist_ok=True)

        new_proto_paths = []
        for path in proto_paths:
            new_path = os.path.join(worker_dir, os.path.basename(path))
            shutil.copy2(path, new_path)
            new_proto_paths.append(new_path)

        args = [
            sys.argv[0],
            f'-I{wd}',
            '--python_out=python',
            '--grpc_python_out=python',
            *new_proto_paths
        ]

        protoc.main(args)
