from setuptools import setup

setup(
    name='ocher',
    version='0.1.0',
    platforms='any',
    packages=['ocher'],
    install_requires=[
        'grpcio >= 1.28.1'
    ],
    extras_require={
        'dev': ['grpcio-tools >= 1.28.1'],
    },
    python_requires='>=3.7',
    author='Max Kuznetsov',
    author_email='maks.kuznetsov@gmail.com',
    description='Python client for Ocher queues',
    license='MIT',
    url='https://github.com/mkuznets/ocher',
    classifiers=[
        'License :: OSI Approved :: MIT License',
        'Operating System :: OS Independent',
        'Programming Language :: Python :: 3 :: Only',
    ],
)
